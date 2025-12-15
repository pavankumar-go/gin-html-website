// let slideIndex = 0;
// showSlides();

// function showSlides() {
//     let i; 
//     let slides = document.getElementsByClassName("slideshow");
//     for (i = 0; i < slides.length; i++) {
//         slides[i].style.display = "none";
//     }
//     slideIndex++;
//     if (slideIndex > slides.length) { slideIndex = 1 }

//     slides[slideIndex - 1].style.display = "block";
//     setTimeout(showSlides, 3000);
// }



let slideIndex = 0;
showSlides();

function showSlides() {
    const slides = document.getElementsByClassName("slideshow");

    for (let i = 0; i < slides.length; i++) {
        slides[i].classList.remove("active");
    }

    slideIndex++;
    if (slideIndex > slides.length) {
        slideIndex = 1;
    }

    slides[slideIndex - 1].classList.add("active");

    setTimeout(showSlides, 3000);
}
