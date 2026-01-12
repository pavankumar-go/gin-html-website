function toggleMenu() {
    const menu = document.getElementById('mobile-menu');
    menu.classList.toggle('active');
}

function showSection(sectionId, isMobile = false) {
    // Hide all sections
    const sections = document.querySelectorAll('section');
    sections.forEach(sec => sec.classList.remove('active'));

    // Remove active class from desktop links
    const desktopLinks = document.querySelectorAll('aside nav a');
    desktopLinks.forEach(link => link.classList.remove('active'));

    // Show target section
    document.getElementById(sectionId).classList.add('active');

    // Highlight specific desktop link
    const desktopLink = document.getElementById('link-' + sectionId);
    if (desktopLink) desktopLink.classList.add('active');

    // Close mobile menu if triggered from mobile
    if (isMobile) {
        toggleMenu();
    }

    // Scroll to top
    window.scrollTo(0, 0);
}


// --- LIGHTBOX FUNCTIONS (NEW) ---
function openLightbox(clickedItem) {
    // Determine which gallery this image belongs to
    const galleryElement = clickedItem.closest('.gallery-grid');
    currentGallery = Array.from(galleryElement.querySelectorAll('.gallery-item'));
    currentIndex = currentGallery.indexOf(clickedItem);

    // const imgSrc = clickedItem.querySelector('img').src;
    const fullResImg = clickedItem.getAttribute('data-full');
    const lightbox = document.getElementById('imageLightbox');
    const lightboxImage = document.getElementById('lightboxImage');

    lightboxImage.src = fullResImg;
    lightbox.style.display = 'block';
    document.body.style.overflow = 'hidden';

    updateNavButtons();
}

function updateNavButtons() {
    const prevBtn = document.getElementById('prevButton');
    const nextBtn = document.getElementById('nextButton');

    // Hide prev button if at the first image
    prevBtn.style.visibility = currentIndex === 0 ? 'hidden' : 'visible';

    // Hide next button if at the last image
    nextBtn.style.visibility = currentIndex === currentGallery.length - 1 ? 'hidden' : 'visible';
}

function navigateLightbox(direction) {
    if (!currentGallery) return;

    if (direction === 'next' && currentIndex < currentGallery.length - 1) {
        currentIndex++;
    } else if (direction === 'prev' && currentIndex > 0) {
        currentIndex--;
    } else {
        return; // Do nothing if at boundary
    }

    const nextImageSrc = currentGallery[currentIndex].querySelector('img').src;
    document.getElementById('lightboxImage').src = nextImageSrc;

    updateNavButtons();
}

function closeLightbox(event) {
    if (event === undefined || event.target.id === 'imageLightbox' || event.target.classList.contains('close-lightbox')) {
        const lightbox = document.getElementById('imageLightbox');
        lightbox.style.display = 'none';
        document.body.style.overflow = 'auto';
        currentGallery = null; // Clear gallery context
        currentIndex = -1;
    }
}
