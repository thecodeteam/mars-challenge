# Setting up the Raspberry Pi Sensors

Logging into to [Raspbian](http://www.raspbian.org/), you will need to use the following credentials:

**user:** pi
**password:** raspbian

### Updating Raspbian

Fist step is to update and upgrade [Raspbian](http://www.raspbian.org/) to the latest version. For this perform the following commands:

    sudo apt-get update
    sudo apt-get upgrade -y


### Installing Go 1.4.2

There are multiple ways of installing Go 1.4.2 on the Raspberry PI. These are two most popular ways we have found:

[http://dave.cheney.net/2012/09/25/installing-go-on-the-raspberry-pi ](http://dave.cheney.net/2012/09/25/installing-go-on-the-raspberry-pi )

and

[https://xivilization.net/~marek/blog/2015/05/04/go-1-dot-4-2-for-raspberry-pi/](https://xivilization.net/~marek/blog/2015/05/04/go-1-dot-4-2-for-raspberry-pi/)

    wget https://xivilization.net/~marek/raspbian/xivilization-raspbian.gpg.key -O - | sudo apt-key add -
    sudo wget https://xivilization.net/~marek/raspbian/xivilization-raspbian.list -O /etc/apt/sources.list.d/xivilization-raspbian.list

    sudo aptitude update
    sudo aptitude install golang


### Download the Mars Sensors' Service

Download the Mar's Sensors' service that will provide you with the Temperature and Radiation time series:

    git clone https://github.com/ghostplant/mars-challenge.git


### Check your Go Environment Variables

Check your Go environment variables by executing `go env`. You will probably have to setup the **GOPATH** variable to point to the folder where you cloned the mars-challenge git repository. Assuming you cloned the repository in the /home/pi folder, you would setup the GOPATH variable like this:

    export GOPATH=/home/pi/mars-challenge/


### Installing the Required Go Packages

The Sensor service requires the following Go packages:

- [Gorilla Websocket](https://github.com/gorilla/websocket)
- [Gorilla Mux](https://github.com/gorilla/mux)

Install using the following Commands:

    go get https://github.com/gorilla/websocket
    go get https://github.com/gorilla/mux


### Executing the Sensors' Service

You are now ready to star the service. Navigate to the `sensor-client` directory in the mars-challenge folder. For example: **/home/pi/mars-challenge/sensor-client**. Then execute the following command:

    cd /home/pi/mars-challenge/sensor-client/
	go run *.go

This will start the service on port 8080. The service will open a Websocket in Port 8080 and also post the weather data on the console.


### Opening Port 8080 on the Raspberry Pi

You may need to open the Raspberry's firewall in order to access the Sensors' service. In order to do so, run the following commands:

    sudo iptables -A INPUT -p tcp --dport 8080 -j ACCEPT
    sudo iptables-save

check that the changes have been persisted:

    sudo iptables -L

### Getting your Sensors' URL

To verify the Sensors' service from another computer in the same network, you can perform the following command:

    sudo ifconfig

The command should provide you with the Raspberry IP address. The port has been setup to `8080`.So, for example the URL could be: http://10.0.0.4:8080
