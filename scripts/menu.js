const openedMenu = document.getElementById('openedMenu');
const rest = document.getElementById('rest');

function openMenu(x) {
    x.classList.toggle("menu-icon-change");
    if (openedMenu.style.display === "none") {
        openedMenu.style.display = "block";
        rest.style.display = 'none';
    } else {
        openedMenu.style.display = "none";
        rest.style.display = 'block';
    }
}

const homeBackground = document.getElementById('bimg');
// for home page
function homeMenu(x) {
    x.classList.toggle("menu-icon-change");
    if (openedMenu.style.display === "none") {
        openedMenu.style.display = "block";
        homeBackground.style.background = "black";
    } else {
        openedMenu.style.display = "none";
        homeBackground.style.background = "";
    }
}