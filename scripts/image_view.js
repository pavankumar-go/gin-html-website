const enlarge = document.querySelectorAll('.image');
const allImages = document.querySelectorAll('.container');
const imageView = document.querySelector('.image-view');
const nextBtn = document.getElementById('next-btn');
const prevBtn = document.getElementById('prev-btn');
const imageBox = document.querySelector('.image-box');

let currentImgIndex = 0;

imageView.addEventListener('click', function () {
    this.style.display = "none";
    imageBox.style.display = "none";
})

enlarge.forEach(function (btn, index) {
    btn.addEventListener('click', function () {
        imageView.style.display = "block";
        imageBox.style.display = "block";
        currentImgIndex = index + 1;
        currentImageDisplay(currentImgIndex);
    })
})

function currentImageDisplay(index) {
    imageBox.style.background = `url(/static/assets/images/${index}.jpg) center/cover no-repeat`
}

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