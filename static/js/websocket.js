var socket;

// message ID
// USER_REGISTER = 0
// USER_LOGIN = 1
// COURT_ADD = 2
// COURT_REMOVE = 3
// COURT_MODIFY = 4
// PLAYER_ADD = 5
// PLAYER_REMOVE = 6

$(document).ready(function () {
    // Create a socket
    socket = new WebSocket('ws://' + window.location.host + '/board/msg?uID=' + $('#uID').text());
    // Message received on the socket
    socket.onmessage = function (event) {
		console.log(event)
        var data = JSON.parse(event.data);
        console.log(data);
        switch (data.Type) {
        case 0: // JOIN
//            if (data.User == $('#uname').text()) {
//                $("#chatbox li").first().before("<li>You joined the chat room.</li>");
//            } else {
//                $("#chatbox li").first().before("<li>" + data.User + " joined the chat room.</li>");
//            }
//			$("#online li").first().before("<li>"+data.User+"</li>");
			$("#online li").first().before("<li>" + data.User + " joined the chat room.</li>");
            break;
        case 1: // LEAVE
//            $("#chatbox li").first().before("<li>" + data.User + " left the chat room.</li>");
            break;
        case 2: // MESSAGE
            $("#chatbox li").first().before("<li><b>" + data.User + "</b>: " + data.Content + "</li>");
            break;
        }
    };

    // Send messages.
    var postConecnt = function () {
        //user ID
		var uID = $('#uID').text();
		
		// week 
		var week = $('#week').val();
		
		// court number
		var courtNum = $('#courtNum').val();
		if (courtNum == ""){
			alert("Court Number Error!")
			return;
		}
		
		// start time 
		var startTime = $('#startTime').val();
		if (startTime == ""){
			alert("Start Time Error!")
			return;
		}
		
		// end time 
		var endTime = $('#endTime').val();
		if (endTime == ""){
			alert("End Time Error!")
			return;
		}
		
		// court type
		var courtType = $('#courtType').val();
		if (courtType == ""){
			alert("Court Type Error!")
			return;
		}
		
		var jsonFormat = {"MsgID":2, "Owner":uID, "Date":week, "Number": courtNum, "Start_time": startTime, "End_time": endTime, "CourtType":courtType}		
		var JSONStr = JSON.stringify(jsonFormat)
		console.log(jsonFormat);
        socket.send(JSONStr);
        $('#publishBtn').val("");
    }

    $('#publishBtn').click(function () {
        postConecnt();
    });
});