const socket = io('ws://localhost:5001', {transports: ['websocket']});
const room = document.getElementById("room");
const form = room.querySelector(".msgForm");
const roomForm = room.querySelector(".roomForm");
const outBtn = room.querySelector(".outBtn");

var roomId = "";
// listen for messages
socket.on('/message', function(message) {

    console.log('new message');
    addMessage(message);
});

function addMessage(msg) {
    const ul = room.querySelector("ul");
    const li = document.createElement("li");
    li.innerText = msg;
    ul.appendChild(li);
}

function handleMsgClick(event) {
    event.preventDefault();
    const input = form.querySelector("input");
    const value = input.value;
    
    socket.emit('/chat', {id: "myID", message: value, room_id: roomId}, function(result) {

        console.log(result);
    });

    input.value = "";
}


function handleRoomClick(event) {
    event.preventDefault();
    const input = roomForm.querySelector("input");
    const value = input.value;
    roomId = value;
    socket.emit('/room', {id: "test", message: "", room_id: value}, function(result) {
        console.log(result);
    });
}


function handleOutClick() {
    socket.emit('/leave', {id: "test", message: "", room_id: roomId}, function(result) {
        console.log(result);
    });
}
form.addEventListener("submit", handleMsgClick);
roomForm.addEventListener("submit", handleRoomClick)
outBtn.addEventListener("click", handleOutClick)