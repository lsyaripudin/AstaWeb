// Preloader
$(window).on('load', function () {
    $('.preloader').fadeOut(800);
});

$(document).ready(function () {
    // Hide mobile menu after clicking on a link
    $('.navbar-collapse a').on('click', function () {
        $('.navbar-collapse').collapse('hide');
    });

    // Navbar shadow on scroll
    $(window).on('scroll', function () {
        if ($('.custom-navbar').offset().top > 40) {
            $('.custom-navbar').addClass('navbar-solid');
        } else {
            $('.custom-navbar').removeClass('navbar-solid');
        }
    });

    // Smooth scroll for anchors with class .smoothScroll
    $('a.smoothScroll').on('click', function (event) {
        var target = this.hash;
        if (target && $(target).length) {
            event.preventDefault();
            $('html, body').animate({
                scrollTop: $(target).offset().top - 60
            }, 700);
        }
    });

    // Initialize wow animations
    if (typeof WOW === 'function') {
        new WOW().init();
    }
});
