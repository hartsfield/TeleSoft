function slideOutAnimation() {
    var nav = document.getElementById("mainNav");
    var interval = setInterval(function() {
        num = parseInt(nav.style.width);
        num -= 3;
        nav.style.width = num + "px";
        if (num <= 0) {
            nav.style.right = "-2em";
            clearInterval(interval);
        }
    }, 1);
    interval();
}

function slideInAnimation() {
    var nav = document.getElementById("mainNav");
    var interval = setInterval(function() {
        num = parseInt(nav.style.width);
        num += 3;
        nav.style.width = num + "px";
        if (num >= 300) {
            nav.style.right = "0px";
            clearInterval(interval);
        }
    }, 1);
    interval();
}
