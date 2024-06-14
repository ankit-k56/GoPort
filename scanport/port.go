package scanport

import (
	"fmt"
	"net"
	"sync"
	"time"
)
type PortStatus struct{
	Port int      `json:"port"`
	Status string  `json:"status"`

}
type Output struct {
	IP string `json:"ip"`
	Protocol string `json:"protocol"`
	OpenPorts []PortStatus `json:"openPorts"`
	ClosedPorts []PortStatus `json:"closedPorts"`
	
}
func TcpScan( port int, host string, wg *sync.WaitGroup, resChan chan<- PortStatus)  {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", address, 3 * time.Second)
	if err != nil{
		resChan <- PortStatus{Port: port, Status: "closed"}
		return
	}
	defer conn.Close()
	resChan<- PortStatus{Port: port, Status: "open"}
	return
}

func UdpScan(port int, host string, wg *sync.WaitGroup, resChan chan<- PortStatus) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", host, port);
	conn, err := net.DialTimeout("udp", address, 1 * time.Second)
	if err != nil{
		resChan<- PortStatus{Port: port, Status: "closed"}
		return
	}
	conn.SetDeadline(time.Now().Add(1 * time.Second))
	_, err = conn.Write([]byte("Hello"))
	if err != nil{
		resChan<- PortStatus{Port: port, Status: "closed"}
		return
	}

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil{
		resChan<- PortStatus{Port: port, Status: "closed"}
		return
	
	}


	defer conn.Close()
	resChan<- PortStatus{Port: port, Status: "open"}


}