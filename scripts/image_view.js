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

// for mobile phones - touch on image to navigate between images
imgTag.addEventListener('click', function () {
    if (currentImgIndex === allImages.length) {
        currentImgIndex = 0;
    }
    currentImgIndex++;
    currentImageDisplay(currentImgIndex);
})

// left arrow - click to go previous image from current image
prevBtn.addEventListener('click', function () {
    if (currentImgIndex === 0) {
        currentImgIndex = allImages.length;
        // disable image view after last image in list is viewed
        navBar.style.display = "block";
        imageView.style.display = "none";
        imageBox.style.display = "none";
    }

    currentImgIndex--;
    currentImageDisplay(currentImgIndex);
})

// right arrow - click to go next image from current image
nextBtn.addEventListener('click', function () {
    currentImgIndex++;
    if (currentImgIndex === allImages.length) {
        currentImgIndex = 0;
        // disable image view after last image in list is viewed
        navBar.style.display = "block";
        imageView.style.display = "none";
        imageBox.style.display = "none";
    }
    currentImageDisplay(currentImgIndex);
})

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
