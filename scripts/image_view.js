const enlarge = document.querySelectorAll('.images');
const allImages = document.querySelectorAll('.container');
const imageView = document.getElementById('image-view');
const mainBody = document.getElementById('main-body');

const nextBtn = document.getElementById('next-btn');
const prevBtn = document.getElementById('prev-btn');
const imageBox = document.querySelector('.image-box');
const imgTag = document.getElementById('image-tag');

// const placeID = document.getElementById('placeID');

let currentImgIndex = 0;

prevBtn.addEventListener('click', function () {
    currentImgIndex--;
    if (currentImgIndex === 0) {
        currentImgIndex = allImages.length;
    }

    currentImageDisplay(currentImgIndex);
})

nextBtn.addEventListener('click', function () {
    currentImgIndex++;
    if (currentImgIndex === allImages.length) {
        currentImgIndex = 1;
    }

    currentImageDisplay(currentImgIndex);
})


enlarge.forEach(function (btn, index) {
    btn.addEventListener('click', function () {
        imageView.style.display = "block";
        imageBox.style.display = "block";
        nextBtn.style.display = "block";
        prevBtn.style.display = "block";
        currentImgIndex = index + 1;
        currentImageDisplay(currentImgIndex);
    })
})

function currentImageDisplay(index) {
    imgTag.src = `/static/assets/images/places/1/${index}.jpg`
    // imageBox.style.background = `url(/static/assets/images/places/1/${index}.jpg) center/cover no-repeat`
}

mainBody.addEventListener('click', function (e) {
    if (e.target.id == "image-view") {
        console.log(e.target.id)
        imageView.style.display = "none";
        imageBox.style.display = "none";
        nextBtn.style.display = "none";
        prevBtn.style.display = "none";
    }
})
