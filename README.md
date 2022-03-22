## Plants Monitoring - Device

Device will capture and send data to server.

### How to works

In the path of project you have to install dependencies

`$ go get ./...`

To be honest I do not use this command.... so if do not works, tell me.

### Building

To build a executable, just follow this commands.

<pre>
	cd .\src\
	go build -o device.exe
</pre>

Or you can use the executable that is at folder "dist".

### Starting the device

To start the device, just follow this commands.

<pre>
	cd .\dist\
	.\device.exe DEADBEEF --interval 5 --api http://localhost:9090
</pre>

The first parameter is the name of device (must to be unique). Other parameters:

- \-\-interval (time that will send data to server, in seconds)
- \-\-api (path to server)
- \-\-user (email of user linked with this device)
- \-\-city (name of city that will capture the data)

Default user is "rbussolo91@gmail.com", has a SQL in backend project with insert of this user.