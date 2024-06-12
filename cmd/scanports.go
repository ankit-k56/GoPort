package cmd

import (
	"strconv"
	"strings"
	"sync"

	"github.com/ankit-k56/GoPort/scanport"
	"github.com/ankit-k56/GoPort/utils"
	"github.com/spf13/cobra"
)



var portStatusesOpen []scanport.PortStatus
var portStatusesClosed []scanport.PortStatus

var scanPort = &cobra.Command{
	Use:   "ping",
	Short: "Used for ping scans",
	Run:   scanPorts,
}

func init() {
	rootCmd.AddCommand(scanPort)
	scanPort.Flags().StringP("port", "p", "", "Port to scan")
	scanPort.Flags().BoolP("udp", "u", false, "Scan using udp")
	scanPort.Flags().StringP("host", "H", "localhost", "Host to scan")
}

func scanPorts(cmd *cobra.Command, args []string){


	portS, _ := cmd.Flags().GetString("port")
	host, _ := cmd.Flags().GetString("host")
	udp, _ := cmd.Flags().GetBool("udp")
	
	
	if portS == ""{
		cmd.Println("Please provide a port to scan using flag p")
		return
	}
	Output := scanport.Output{};
	Output.IP = host
	if(udp){
		Output.Protocol = "udp"
	}else{
		Output.Protocol = "tcp"
	}

	wg := sync.WaitGroup{}
	resChan := make(chan scanport.PortStatus)


	
	portsArrayComma := strings.Split(portS, ",")
	if(len(portsArrayComma) >= 2){
		for _ ,port := range portsArrayComma {
			port, err := strconv.Atoi(port)
			if err != nil{
				cmd.Println("Error in converting port to int , no comma separated ports found")
				continue
			}
			scanport.PingScan(port, host, &wg, resChan)			
		}
	}
		
	portsArrayHyphen := strings.Split(portS, "-")
	if len(portsArrayHyphen) == 2{
		startPort, err := strconv.Atoi(portsArrayHyphen[0])
		if err != nil{
			cmd.Println("Error in converting port to int, no startPort found")
			return
		}
		endPort, err := strconv.Atoi(portsArrayHyphen[1])
		if err != nil{
			cmd.Println("Error in converting port to int, no endPort found")
			return
		}
		for i := startPort; i<= endPort; i++{
			wg.Add(1)
			if udp{
				go scanport.UdpScan(i, host, &wg, resChan)
	
			}else{
				go scanport.PingScan(i, host, &wg, resChan)
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
		cmd.Println("Scan completed, Check /Output/output.json for results")

}
