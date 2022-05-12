package handlers

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/nyakaspeter/raven-torrent/internal/torrentclient"
)

var webSocket *websocket.Conn
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ReceiveTorrent() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		io.WriteString(w, torrentclient.CheckReceiver(vars["todo"]))
	}
}

func ReceiverPage(version string, apiPrefix string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, createReceiverPage(version, apiPrefix))
	}
}

func Websocket(procQuit chan bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		webSocket, _ = upgrader.Upgrade(w, r, nil) // Error ignored

		for {
			// Read message from ws
			messageType, message, err := webSocket.ReadMessage()
			if err != nil {
				return
			}

			if messageType == 1 {
				if string(message) == "stop" {
					if err = webSocket.WriteMessage(1, []byte("{\"function\":\"stopserver\",\"data\": \"ok\"}")); err != nil {
						return
					}

					go func() {
						time.Sleep(1 * time.Nanosecond)
						procQuit <- true
					}()
				} else {
					value := torrentclient.SetReceivedMagnet(string(message))
					if err = webSocket.WriteMessage(1, []byte("{\"function\":\"sendmagnet\",\"data\":\""+value+"\"}")); err != nil {
						return
					}
				}
			} else if messageType == 2 {
				metaData, error := metainfo.Load(bytes.NewReader(message))
				if error == nil {
					value := torrentclient.SetReceivedTorrent(metaData)
					if err = webSocket.WriteMessage(1, []byte("{\"function\":\"sendfile\",\"data\":\""+value+"\"}")); err != nil {
						return
					}
				}
			}

		}
	}
}

func createReceiverPage(version string, apiPrefix string) string {
	html := `<!DOCTYPE html>
			<html lang="en">
			<head>
			  	<meta charset="UTF-8">
			  	<meta name="viewport" content="width=device-width, initial-scale=1.0">
			  	<title>White Raven Server v` + version + `</title>
			  	<link rel="icon" type="image/png" sizes="32x32" href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAgY0hSTQAAeiYAAICEAAD6AAAAgOgAAHUwAADqYAAAOpgAABdwnLpRPAAAAAZiS0dEAAAAAAAA+UO7fwAAAAlwSFlzAAAASAAAAEgARslrPgAAAmdJREFUWMPt10uIz1EUB/APozQL7CwU5bUUeSZm4TmxmEYhKxsLK8qGUvhPFCMLZMNapryZkjxWxisprLwXClFE0qDGWNzfz//69X/1f5HmW7d+99zzu+f7O+fcc+6PIQyhNgzLzNsxDS1Nsj+AR7hSaPEEBv/SOJ4l0/4XyaRjKYxICM2MyN1EX5NCthALkuc5uJYSGhkpXUOuSYRyEaGRMLxJhivGiDLrsUvrhZIpUY7QUuyqM6GuSgjlFM6bwTqTKYtyObQXa3H1XyH0A6ewHG241SxCOXzGS1zCTszL6PYJSb4F35rhodGYiBVC4t3BU2xCa6IziIOYjzeNJlQIU3EYT9AZyR8kpB43itBujMMsrBfy5kukNx7ncFS+qr8SemBDPFUIrdiKj/5sgtcxKtKbgX6VN9Fc9G4uKy8Vsn7sxxScjuSLcSHy1ANsr5cXUkI7BNffxR5Mj3Q+CrWoS75QLhJyK8Uh3K4Xqazr0nFWSOpSep3RWpsGh2wV7mFlJOvCmYxn0pJwQ8ivmpA9ZR3owc9EPgYXhdok+ZKN+JTMJ2BDtN+xWgmlzXUAb9GbjAPC0Z8kXPh7hJLwAh+wLxmwGUeS5/OJfFgJm1XfRscKrSSNcRyqVqFOpWtzq7SRUySHduC5kCNpTrzHGvnwrZI/ff24HG28vNqvziINWQsmC011GZYkRu/jJNYJYViNh8k7vcmc0EZibKvQfnwb/V7Mddmj2RHJ41ozO5I/zxiq5jdoWSlCz6K1cZH8dRH55xoJ/f5RTEPWh+5ow4Ho+V20Vkz+05/oVhm+Ctecpt1Ih/D/4xfB5AZs0Y/GewAAACV0RVh0ZGF0ZTpjcmVhdGUAMjAxOS0xMC0yMVQxNjo1MzoyNSswMDowMHy2Ji8AAAAldEVYdGRhdGU6bW9kaWZ5ADIwMTktMTAtMjFUMTY6NTM6MjUrMDA6MDAN656TAAAAKHRFWHRzdmc6YmFzZS11cmkAZmlsZTovLy90bXAvbWFnaWNrLVBkYXlnNXRhwVfHAwAAAABJRU5ErkJggg==">
			  	<script src="http://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
			  	<script>
	               	$(document).ready(function () {
	               		var socket = new WebSocket('ws://' + window.location.host + '` + apiPrefix + `websocket');

	                    $('#sendbutton').click(function () {
	                    	if (document.getElementById('addmagnet').parentElement.className.indexOf("focus") != -1) {
		                    	var magnetlink = document.getElementById("magnetlink").value;
		                    	var result = magnetlink.match(/magnet:\?xt=urn:btih:([a-zA-Z0-9]*)/);

		                    	SetInputsDisabledState(true);

		                    	if (result && (result[1].length == 40 || result[1].length == 32)) {
		                    		var infodiv = document.getElementById("info");
		                    		infodiv.className = "alert alert-info";
			                        infodiv.innerHTML = "<strong>WAIT!</strong> TRYING TO SEND MAGNET LINK...";
			                        infodiv.style.display = "block";

			                        if (socket.readyState == 1) {
			                        	socket.send(magnetlink);
			                        } else {
			                        	infodiv.className = "alert alert-danger";
                        				infodiv.innerHTML = "<strong>ERROR!</strong> UNABLE TO ADD MAGNET LINK!";
                        				infodiv.style.display = "block";

		                        		document.getElementById("magnetlink").value = "";

		                        		SetInputsDisabledState(false);

		                        		setTimeout(function(){
		                        			infodiv.style.display = "none";
		                        		}, 3000);
			                        }
			                    } else {
			                    	if (magnetlink != "") {
			                    		var infodiv = document.getElementById("info");
			                    		infodiv.className = "alert alert-warning";
			                        	infodiv.innerHTML = "<strong>WARNING!</strong> THIS MAGNET LINK IS NOT VALID!";
			                        	infodiv.style.display = "block";
			                        	setTimeout(function(){ infodiv.style.display = "none"; }, 3000);
			                    	}
			                    	SetInputsDisabledState(false);
			                    }
			                } else {
			                	var filelink = document.getElementById('filelink');
			                	if (filelink.files[0]) {
			                		SetInputsDisabledState(true);
																						                	
				                	var infodiv = document.getElementById("info");
		                    		infodiv.className = "alert alert-info";
			                        infodiv.innerHTML = "<strong>WAIT!</strong> TRYING TO ADD TORRENT FILE...";
			                        infodiv.style.display = "block";

			                        if (socket.readyState == 1) {
			                        	socket.send(filelink.files[0]);
			                        } else {
			                        	infodiv.className = "alert alert-danger";
                        				infodiv.innerHTML = "<strong>ERROR!</strong> UNABLE TO ADD TORRENT FILE!";
                        				infodiv.style.display = "block";

		                        		var filename = document.getElementById('filename');
		                        		filename.innerText = "Stranger.Things.S03E08.480p.x264-mSD[eztv].mkv.torrent"
		                				filename.style.color = "#acb6c0";
		                				document.getElementById('filelink').value = "";

		                        		SetInputsDisabledState(false);

		                        		setTimeout(function(){
		                        			infodiv.style.display = "none";
		                        		}, 3000);
			                        }
				                }
			                }
	                    });

	                    $('#power').click(function () {
	                    	document.getElementsByClassName('container')[0].style.display = 'none';
	                        document.getElementById('power').style.display = 'none';
	                        document.getElementById('h1t').innerText = 'BYE BYE!';
	                        document.getElementById('h4t').innerText = 'TRYING TO STOP THE SERVER!';

	                    	if (socket.readyState == 1) {
	                    		socket.send('stop');
	                    	} else {
	                    		document.getElementById('h4t').innerText = 'WHITE RAVEN SERVER ALREADY STOPPED!';
	                    	}
	                    });

	                    socket.onmessage = function (e) {
	                    	var response = JSON.parse(e.data);
	                    	var infodiv = document.getElementById("info");

	                    	if (response.function == "stopserver") {
	                    		if (response.data == "ok") {
	                    			document.getElementById('h4t').innerText = 'WHITE RAVEN SERVER STOPPED!';
	                    		} else {
	                    			document.getElementById('h4t').innerText = 'WHITE RAVEN SERVER ALREADY STOPPED!';
	                    		}
	                    	} else if (response.function == "sendmagnet") {
	                    		if (response.data != "") {
                        			infodiv.className = "alert alert-success";
                        			infodiv.innerHTML = "<strong>SUCCESS!</strong> TORRENT ADDED SUCCESSFULLY!";
                        		} else {
                        			infodiv.className = "alert alert-danger";
                        			infodiv.innerHTML = "<strong>ERROR!</strong> UNABLE TO ADD MAGNET LINK!";
                        		}
                        		infodiv.style.display = "block";

                        		document.getElementById("magnetlink").value = "";

                        		SetInputsDisabledState(false);

                        		setTimeout(function(){
                        			infodiv.style.display = "none";
                        		}, 3000);
	                    	} else if (response.function == "sendfile") {
	                    		if (response.data != "") {
                        			infodiv.className = "alert alert-success";
                        			infodiv.innerHTML = "<strong>SUCCESS!</strong> TORRENT ADDED SUCCESSFULLY!";
                        		} else {
                        			infodiv.className = "alert alert-danger";
                        			infodiv.innerHTML = "<strong>ERROR!</strong> UNABLE TO ADD TORRENT FILE!";
                        		}
                        		infodiv.style.display = "block";

                        		var filename = document.getElementById('filename');
                        		filename.innerText = "Stranger.Things.S03E08.480p.x264-mSD[eztv].mkv.torrent"
                				filename.style.color = "#acb6c0";
                				document.getElementById('filelink').value = "";

                        		SetInputsDisabledState(false);

                        		setTimeout(function(){
                        			infodiv.style.display = "none";
                        		}, 3000);
	                    	}	                        
	                    };

	                    socket.onclose = function (e) {
	                    	document.getElementsByClassName('container')[0].style.display = 'none';
	                        document.getElementById('power').style.display = 'none';
	                        document.getElementById('h1t').innerText = 'BYE BYE!';
	                    	document.getElementById('h4t').innerText = 'WHITE RAVEN SERVER STOPPED!';
	                    }
	                });

	                function SelectMagnet() {
                    	document.getElementById('addfile').parentElement.className = "btn btn-primary"
                    	document.getElementById('addmagnet').parentElement.className = "btn btn-primary active focus";
                    	document.getElementById('fileupload').style.display = "none";
                    	document.getElementById('magnetlink').style.display = "block";
                    }

                    function SelectFile() {
                    	document.getElementById('addmagnet').parentElement.className = "btn btn-primary"
                    	document.getElementById('addfile').parentElement.className = "btn btn-primary active focus";
                    	document.getElementById('magnetlink').style.display = "none";
                    	document.getElementById('fileupload').style.display = "block";
                    }

                    function CheckFileData() {
                    	var filename = document.getElementById('filename');
                    	if (event.target.files[0]) {
                    		var fileext = event.target.files[0].name.match(/(\.torrent)$/i);
	                    	if (event.target.files[0].type == "application/x-bittorrent" || fileext != null) {
	                    		if (event.target.files[0].size < 524288) {
			                    	filename.innerText = event.target.files[0].name;
			                    	filename.style.color = "#000000";
			                    	//event.target.value = '';
			                    } else {
			                    	var infodiv = document.getElementById("info");
		                    		infodiv.className = "alert alert-warning";
		                        	infodiv.innerHTML = "<strong>WARNING!</strong> TORRENT FILE SIZE IS TO BIG. MAX ALLOWED SIZE: 0.5 MB";
		                        	infodiv.style.display = "block";
		                        	event.target.value = '';
		                        	filename.innerText = "Stranger.Things.S03E08.480p.x264-mSD[eztv].mkv.torrent"
			                    	filename.style.color = "#acb6c0";
		                        	setTimeout(function(){ infodiv.style.display = "none"; }, 3000);
			                    }
		                    } else {
	                    		var infodiv = document.getElementById("info");
	                    		infodiv.className = "alert alert-warning";
	                        	infodiv.innerHTML = "<strong>WARNING!</strong> WRONG FILE EXTENSION DETECTED.";
	                        	infodiv.style.display = "block";
	                        	event.target.value = '';
	                        	filename.innerText = "Stranger.Things.S03E08.480p.x264-mSD[eztv].mkv.torrent"
			                    filename.style.color = "#acb6c0";
	                        	setTimeout(function(){ infodiv.style.display = "none"; }, 3000);
		                    }
		                } else {
		                    filename.innerText = "Stranger.Things.S03E08.480p.x264-mSD[eztv].mkv.torrent"
			                filename.style.color = "#acb6c0";
		                }
                    }

                    function SetInputsDisabledState(state) {
                    	document.getElementById('sendbutton').disabled = state;
                    	var inputs = document.getElementsByTagName('input');
						for(var i = 0; i < inputs.length; i++) {
						    inputs[i].disabled = state;
						    if (inputs[i].type == 'radio') {
						    	if (state) {
						    		inputs[i].parentElement.style.cursor = "not-allowed";
						    	} else {
						    		inputs[i].parentElement.style.cursor = "pointer";
						    	}
						    } else {
						    	if (state) {
						    		inputs[i].style.cursor = "not-allowed";
						    	} else {
						    		inputs[i].style.cursor = "default";
						    	}
						    }
						}
						if (state) {
							document.getElementById('filename').style.cursor = "not-allowed";
						} else {
							document.getElementById('filename').style.cursor = "pointer";
						}
                    }
	            </script>

			  	<link rel='stylesheet' href='http://maxcdn.bootstrapcdn.com/bootswatch/3.4.1/flatly/bootstrap.min.css'>
			  	<style type="text/css">
			  		@media (max-width: 710px) {
					    .heading {
					    	max-width: 500px;
    						margin: auto; 
					    }
					}

				  	body {
					  background-color: #d0d0d0;
					  padding: 3%; /*3.125em*/
					  min-width: 500px;
					}

					.container {
					  padding: 0px 20px 0px 20px;
					  background-color: #fff;
					  border-radius: 8px;
					  max-width: 800px;
					}

					.heading {
					  text-align: center;
					}
					.heading h1 {
					  text-align: center;
					  margin: 0 0 5px 0;
					  font-weight: 900;
					  font-size: 4rem;
					  color: #000000;
					}
					.heading h4 {
					  color: #000000;
					  text-align: center;
					  margin: 0 0 35px 0;
					  font-weight: 400;
					  font-size: 24px;
					}

					.btn {
					  outline: none !important;
					}

					.btn.btn-primary {
					  background-color: #383838;
					  border-color: #383838;
					  outline: none;
					}
					.btn.btn-primary:hover {
					  background-color: #505050;
					  border-color: #505050;
					}
					.btn.btn-primary:active {
					  background-color: #383838;
					  border-color: #383838;
					}
					.btn.btn-primary .fa {
					  padding-right: 4px;
					}

					.form-group {
					  margin-top: 20px;
					  margin-bottom: 20px;
					  text-align: center;
					}

					.form-control {
					  text-align: center;
					}

					.alert {
					  margin-top: 20px;
					  border-radius: 4px;
					  text-align: center;
					}
					
					#info {
					  display: none;
					}

					.btn-group {
						margin-bottom: 20px;
					}

					.btn-primary:active:hover,
					.btn-primary.active:hover,
					.open>.dropdown-toggle.btn-primary:hover,
					.btn-primary:active:focus,
					.btn-primary.active:focus,
					.open>.dropdown-toggle.btn-primary:focus,
					.btn-primary:active.focus,
					.btn-primary.active.focus,
					.open>.dropdown-toggle.btn-primary.focus {
						color: #ffffff;
						background-color: #196eab;
						border-color: #196eab;
					}

					#filelink, #fileupload {
						display: none;
					}

					#filename {
						color: #acb6c0;
						cursor: pointer;
						height: 100%;
						overflow: hidden;
						word-wrap: break-word;
					}

					#magnetlink {
						color: #000000;
						font-weight: bold;
					}

					#magnetlink:-ms-input-placeholder {
						color: #acb6c0;
					}

					#website {
						background-color: #196eab;
						color: white;
						transform: rotateZ(-45deg);
						-webkit-transform: rotateZ(-45deg);
						width: 250px;
						position: absolute;
						left: -70px;
						top: 35px;
						font-size: 13px;
						padding: 10px 0px;
						text-align: center;
						box-shadow: 0px 0px 20px #333;
					}

					#website:hover {
						background-color: #397bab;
					}

					#power {
						position: absolute;
						top: 8px;
						left: 8px;
						width: 30px;
						height: 30px;
						background-color: #dc0000;
						border-radius: 15px;
						cursor: pointer;
						box-shadow: 0px 0px 20px #333;
					}

					#power:hover {
						background-color: #fc0000;
					}

			  	</style>
			</head>
			<body>
				<div class="heading">
					<h1 id="h1t">ADD TORRENT MANUALLY</h1></br>
					<h4 id="h4t">INSERT A MAGNET LINK OR UPLOAD A TORRENT FILE</h4>
				</div>

				<div class="container">
					<div id="info" class="alert">
  					</div>	
					<div class="form-group">
						<div class="btn-group btn-group-toggle" data-toggle="buttons">
						  <label class="btn btn-primary active focus">
						    <input type="radio" name="torrent" id="addmagnet" autocomplete="off" onclick="SelectMagnet()"> MAGNET LINK
						  </label>
						  <label class="btn btn-primary">
						    <input type="radio" name="torrent" id="addfile" autocomplete="off" onclick="SelectFile()"> TORRENT FILE
						  </label>
						</div>
						</br>
						<label id="fileupload" for="filelink">
							<span id="filename" class="form-control">Stranger.Things.S03E08.480p.x264-mSD[eztv].mkv.torrent</span>
							<input id="filelink" type="file" accept=".torrent" class="form-control" onchange="CheckFileData()"/>
						</label>
						<input id="magnetlink" type="text" class="form-control" placeholder="magnet:?xt=urn:btih:13938f71a22c4fb4efe112ba76a343a9ea7b33cc"/>
					</div>
					<div class="form-group">
				    	<button id="sendbutton" type="button" class="btn btn-primary btn-block">SEND TO WHITE RAVEN</button>
					</div>
				</div>
				<a href="http://github.com/silentmurdock"><div id="website">White Raven's Website</div></a>
				<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAgY0hSTQAAeiYAAICEAAD6AAAAgOgAAHUwAADqYAAAOpgAABdwnLpRPAAAAAZiS0dEAAAAAAAA+UO7fwAAAAlwSFlzAAAASAAAAEgARslrPgAAAg1JREFUWMPtlz1rFFEUht+j0So2IQF1k6CNELBMp6Cx1UQhMaZIKWJj+nTp1NJOURvxFwjBxiAu2SI/QQMpdN1IsElK4+5jkbvkMDA7XxcCMi8MnLkf7zznzL13GKlWrWqyqgbAdUnXwm3LzDZPNCNgjWOtVfU7daLZ/FdAwAKwWvQBwCowH5UaeAz0wjp5luhLXUOurwesxIJZcjAA34HzWUDAReCn6+sB97OeN/CVAROSXuv4eGhJmjazX1nGZtaRNC1pq98k6S0wXhpI0nNJwyFuS7pnZnt5q2tmu5LmJHVC0zlJT0sBAZckLbqmZTP7nRfGQe1JWnZNS8BkmQotSjod4i0z+1IUxkF9lrQRbocSieYGuuXiV2VhnN6neOcGuuriZgSgVop3bqBRF7cjAP1w8VgZoEMXn4kAdNbFf9IGDQ0w6Ei6EuKGpK8p4zZ1dDz04zQ1XLxbOB3goztlH1YtD/DI+a2njRv0yvykB1WBEh7rhWcDE8Bfl9WNCtWZcT6HWZ+PQUZvnNEOMFrCYwTYdj4vyyYmoAEcOLMmMFIQpunm7wMXSgMF0zmg60y/ATdzzJtJVKYLzFaCceYrCSiAT2HnTAHD4ZoKbRuJsV3gSRSYRKUOKK594E5UGAc1BrwIOyVLXeBd0TVT6kcxbNu7km5Luiypv43bknZ0dM58MLMY38BatWpF1T8njjLwLgYRQgAAACV0RVh0ZGF0ZTpjcmVhdGUAMjAxOS0xMC0yN1QwNjo1NzoyNCswMDowMB20BXMAAAAldEVYdGRhdGU6bW9kaWZ5ADIwMTktMTAtMjdUMDY6NTc6MjQrMDA6MDBs6b3PAAAAKHRFWHRzdmc6YmFzZS11cmkAZmlsZTovLy90bXAvbWFnaWNrLU4tZUZjUjlpSONKeQAAAABJRU5ErkJggg==" title="Stop The Server" id="power">
			</body>
			</html>`

	return html
}
