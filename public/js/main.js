// function slideOutAnimation() {
//     var nav = document.getElementById("mainNav");
//     var interval = setInterval(function() {
//         num = parseInt(nav.style.width);
//         num -= 3;
//         nav.style.width = num + "px";
//         if (num <= 0) {
//             nav.style.right = "-2em";
//             clearInterval(interval);
//         }
//     }, 1);
//     interval();
// }

// function slideInAnimation() {
//     var nav = document.getElementById("mainNav");
//     var interval = setInterval(function() {
//         num = parseInt(nav.style.width);
//         num += 3;
//         nav.style.width = num + "px";
//         if (num >= 300) {
//             nav.style.right = "0px";
//             clearInterval(interval);
//         }
//     }, 1);
//     interval();
// }
//

function sendMail() {
    document.getElementById("sendButt").innerHTML = "Sending...";
    document.getElementById("sendButt").style.backgroundColor = "yellow";
    document.getElementById("sendButt").style.color = "black";
    var xhr = new XMLHttpRequest();

    xhr.open('POST', '/sendMail');
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onload = function() {
        if (xhr.status === 200) {
            var res = JSON.parse(xhr.responseText);
            if (res.success == "true") {
                document.getElementById("sendButt").innerHTML = "SENT!";
                document.getElementById("sendButt").style.backgroundColor = "green";
                document.getElementById("sendButt").style.color = "white";
            } else {
                document.getElementById("sendButt").innerHTML = "Error";
            }
        }
    };

    xhr.send(JSON.stringify({
        from: document.getElementById("emailInput").value,
        body: document.getElementById("messageArea").value
    }));

}
