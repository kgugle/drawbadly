var canvas, ctx, flag = false,
    prevX = 0,
    currX = 0,
    prevY = 0,
    currY = 0;
var penColor = "black",
    penWidth = 2;

var gameSocket = null;

function init() {
    gameSocket = new WebSocket("ws://localhost:9000/game");

    gameSocket.onmessage = function (event) {
      drawReceivedPixel(event.data);
    }

    canvas = document.getElementById('can');
    ctx = canvas.getContext("2d");
    w = canvas.width;
    h = canvas.height;

    canvas.addEventListener("mousemove", function (e) {
        findxy('move', e)
    }, false);
    canvas.addEventListener("mousedown", function (e) {
        findxy('down', e)
    }, false);
    canvas.addEventListener("mouseup", function (e) {
        findxy('up', e)
    }, false);
    canvas.addEventListener("mouseout", function (e) {
        findxy('out', e)
    }, false);        
}

function sendPixel(currX, currY) {
    var pixelBuffer = new Uint32Array(2);
    pixelBuffer[0] = currX;
    pixelBuffer[1] = currY;
    gameSocket.send(pixelBuffer.buffer);
}

function drawReceivedPixel(pixel_data) {
  var res = pixel_data.split(" "),
      x = res[0],
      y = res[1];
  ctx.fillStyle = penColor;
  ctx.fillRect(x, y, penWidth, penWidth);
}

function color(obj) {
    switch (obj.id) {
        case "green":
            penColor = "green";
            break;
        case "blue":
            penColor = "blue";
            break;
        case "red":
            penColor = "red";
            break;
        case "yellow":
            penColor = "yellow";
            break;
        case "orange":
            penColor = "orange";
            break;
        case "black":
            penColor = "black";
            break;
        case "white":
            penColor = "white";
            break;
    }
    if (penColor == "white") penWidth = 14;
    else penWidth = 2;

}

function draw() {
    ctx.beginPath();
    ctx.moveTo(prevX, prevY);
    ctx.lineTo(currX, currY);
    ctx.strokeStyle = penColor;
    ctx.lineWidth = penWidth;
    ctx.stroke();
    ctx.closePath();
}

function erase() {
    var m = confirm("Want to clear");
    if (m) {
        ctx.clearRect(0, 0, w, h);
        document.getElementById("canvasimg").style.display = "none";
    }
}

function save() {
    document.getElementById("canvasimg").style.border = "2px solid";
    var dataURL = canvas.toDataURL();
    document.getElementById("canvasimg").src = dataURL;
    document.getElementById("canvasimg").style.display = "inline";
}

function findxy(res, e) {
    if (res == 'down') {
        prevX = currX;
        prevY = currY;
        currX = e.clientX - canvas.offsetLeft;
        currY = e.clientY - canvas.offsetTop;

        flag = true;
        dot_flag = true;
        if (dot_flag) {
            ctx.beginPath();
            ctx.fillStyle = penColor;
            ctx.fillRect(currX, currY, penWidth, penWidth);
            ctx.closePath();
            dot_flag = false;

            sendPixel(currX, currY);
        }
    }
    if (res == 'up' || res == "out") {
        flag = false;
    }
    if (res == 'move') {
        if (flag) {
            prevX = currX;
            prevY = currY;
            currX = e.clientX - canvas.offsetLeft;
            currY = e.clientY - canvas.offsetTop;
            draw();

            sendPixel(currX, currY);
        }
    }
}
