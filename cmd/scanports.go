package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/ankit-k56/GoPort/scanport"
	"github.com/ankit-k56/GoPort/utils"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)


const maxCurrentScans = 20

var portStatusesOpen []scanport.PortStatus
var portStatusesClosed []scanport.PortStatus

var scanPort = &cobra.Command{
	Use:   "ping",
	Short: "Used for ping scans",
	Run:   scanPorts,
}

func init() {
	rootCmd.AddCommand(scanPort)
	scanPort.SetHelpFunc(customHelpFuncPing)
	scanPort.Flags().StringP("port", "p", "8090", "Port to scan")
	scanPort.Flags().BoolP("udp", "u", false, "Scan using udp")
	scanPort.Flags().StringP("host", "a", "localhost", "Host to scan")
}

func scanPorts(cmd *cobra.Command, args []string){


	portS, _ := cmd.Flags().GetString("port")
	host, _ := cmd.Flags().GetString("host")
	udp, _ := cmd.Flags().GetBool("udp")
	
	

	Output := scanport.Output{};
	Output.IP = host
	if(udp){
		Output.Protocol = "udp"
	}else{
		Output.Protocol = "tcp"
	}

	wg := sync.WaitGroup{}
	resChan := make(chan scanport.PortStatus)
	semaphore := make(chan struct{}, maxCurrentScans)

	singlePort, err := strconv.Atoi(portS)
	if err == nil{
		fmt.Println(yellow("Scanning ..."))
		wg.Add(1)
		semaphore <- struct{}{}
		if udp{
			go scanport.UdpScan(singlePort, host, &wg, resChan)
			<-semaphore

		}else{
			go scanport.TcpScan(singlePort, host, &wg, resChan)
			<-semaphore
		}
		

	}
	
	portsArrayComma := strings.Split(portS, ",")
	if(len(portsArrayComma) >= 2){
		fmt.Println(yellow("Scanning ..."))
		for _ ,port := range portsArrayComma {
			port, err := strconv.Atoi(port)
			
			if err != nil{
				cmd.Println(red("Error in converting port to int , no comma separated ports found"))
				continue
			}
			wg.Add(1)
			semaphore <- struct{}{}
			if udp{
				go scanport.UdpScan(port, host, &wg, resChan)
				<-semaphore
			}else{

				go scanport.TcpScan(port, host, &wg, resChan)	
				<-semaphore		
			}
		}
	}
		
	portsArrayHyphen := strings.Split(portS, "-")
	if len(portsArrayHyphen) == 2{
		fmt.Println(yellow("Scanning ..."))
		startPort, err := strconv.Atoi(portsArrayHyphen[0])
		if err != nil{
			cmd.Println(red("Error in converting port to int, no Start Port found"))
			return
		}
		endPort, err := strconv.Atoi(portsArrayHyphen[1])
		if err != nil{
			cmd.Println(red("Error in converting port to int, no End Port found"))
			return
		}
		for i := startPort; i<= endPort; i++{

			wg.Add(1)
			if udp{
				go scanport.UdpScan(i, host, &wg, resChan)
	
			}else{
				go scanport.TcpScan(i, host, &wg, resChan)
				}
			}
		}

		go func(){
			wg.Wait()
			close(resChan)
		}()
		for result := range resChan{
			if(result.Status == "open"){
				portStatusesOpen = append(portStatusesOpen, result)
			}else{
				portStatusesClosed = append(portStatusesClosed, result)
			}
	
		}
		Output.OpenPorts = portStatusesOpen
		Output.ClosedPorts = portStatusesClosed
		utils.GenerateOutput(Output)
		cmd.Println( green("Scan completed, Check /Output/output.json for results"))

}


func customHelpFuncPing(cmd *cobra.Command, args []string) {
	// tile := figure.NewFigure("GoPort", "", true)

	// fmt.Println(green(tile.String()))
	// fmt.Println(green("A simple and Fast port scanning tool written in GoLang"))
	
	fmt.Println("Usage:")
	fmt.Println(green("  goport ping [flags]"))

	fmt.Println()
	
	fmt.Println("Available Commands:")
	fmt.Println(yellow("  help        Help about any command"))
	fmt.Println()
	fmt.Println("Flags:")
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		fmt.Println(termenv.String(fmt.Sprintf(green("  -%s, --%s\t%s"), flag.Shorthand, flag.Name, flag.Usage)))
	})
	fmt.Println()
	fmt.Println(termenv.String(green(`Use "goport [command] --help" for more information about a command.`)))
}
