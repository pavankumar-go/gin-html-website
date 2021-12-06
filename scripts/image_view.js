const enlarge = document.querySelectorAll('.images');
const allImages = document.querySelectorAll('.container');

const imageView = document.getElementById('image-view');
const mainBody = document.getElementById('main-body');

const nextBtn = document.getElementById('next-btn');
const prevBtn = document.getElementById('prev-btn');

const imageBox = document.querySelector('.image-box');
const imgBox = document.getElementById('img-box');
const imgTag = document.getElementById('image-tag');

const navBar = document.querySelector('.nav-bar');

var imgs = document.getElementsByClassName('imgs');

let currentImgIndex = 0;


var midPoint;
var leftThreshold;
var rightThreshold;
var touchAt;
// ------------------------------------ 
// 0                                355
//                237
// 0       x      237         x     255

// for swiping images
function touchStart(event) {
    touchAt = event.touches[0].clientX
    midPoint = imgTag.clientWidth / 2;
    const threshold = midPoint / 4;
    rightThreshold = midPoint + threshold;
    leftThreshold = midPoint - threshold;
    imgTag.
    imgTag.style.animation = 'splash 1s normal forwards ease-in-out'
}

// when touch/swipe ended, determine the direction of swipe
function touchEnd() {
    if (touchAt > midPoint && touchAt >= rightThreshold) {
        displayNextImg()
    } else if (touchAt < midPoint && touchAt <= leftThreshold) {
        displayPrevImg()
    }
}

// left arrow - click to go previous image from current image
prevBtn.addEventListener('click', function () {
    displayPrevImg()
})

function displayPrevImg() {
    if (currentImgIndex === 0) {
        currentImgIndex = allImages.length;
        // disable image view after last image in list is viewed
        navBar.style.display = "block";
        imageView.style.display = "none";
        imageBox.style.display = "none";
    }

    currentImgIndex--;
    currentImageDisplay(currentImgIndex);
}

// right arrow - click to go next image from current image
nextBtn.addEventListener('click', function () {
    displayNextImg()
})

function displayNextImg() {
    currentImgIndex++;
    if (currentImgIndex === allImages.length) {
        currentImgIndex = 0;
        // disable image view after last image in list is viewed
        navBar.style.display = "block";
        imageView.style.display = "none";
        imageBox.style.display = "none";
    }
    currentImageDisplay(currentImgIndex);
}

// get all images and view individual on click
enlarge.forEach(function (btn, index) {
    btn.addEventListener('click', function () {
        navBar.style.display = "none";
        imageView.style.display = "block";
        imageBox.style.display = "block";
        currentImgIndex = index;

        // disable arrow keys for screens less than 820px (ipad onwards)
        if (window.screen.width <= 820) {
            var deviceWidth = window.screen.width - 20;
            // set width to device screen size - 20px
            imgBox.style.width = deviceWidth + "px";
            nextBtn.style.display = "none";
            prevBtn.style.display = "none";
        } else {
            nextBtn.style.display = "block";
            prevBtn.style.display = "block";
        }

        currentImageDisplay(currentImgIndex);
    })
})

function currentImageDisplay(index) {
    // added for mobile click - it starts with 1 .. for some reason - 1st missed was getting missed
    if (currentImgIndex === allImages.length) {
        imgTag.src = imgs[0].src;
        // disable image view after last image in list is viewed
        navBar.style.display = "block";
        imageView.style.display = "none";
        imageBox.style.display = "none";
    }
    imgTag.src = imgs[index].src;
}

// exist single image view by clicking anywhere other than the image or arrow keys
// TODO: add X icon if needed.
mainBody.addEventListener('click', function (e) {
    if (e.target.id == "image-view") {
        imageView.style.display = "none";
        imageBox.style.display = "none";
        nextBtn.style.display = "none";
        prevBtn.style.display = "none";
        navBar.style.display = "block";
    }
})
