// Preloader
$(window).on('load', function () {
    $('.preloader').fadeOut(1000);
});

$(document).ready(function () {
    // Hide mobile menu after clicking on a link
    $('.navbar-collapse a').on('click', function () {
        $('.navbar-collapse').collapse('hide');
    });

    // Navbar style on scroll
    $(window).on('scroll', function () {
        if ($('.navbar').offset().top > 50) {
            $('.navbar-fixed-top').addClass('top-nav-collapse');
        } else {
            $('.navbar-fixed-top').removeClass('top-nav-collapse');
        }
    });

    // FlexSlider hero
    if ($('.flexslider').length) {
        $('.flexslider').flexslider({
            animation: 'fade',
            directionNav: false,
        });
    }

    // Isotope filtering
    if ($('.iso-box-wrapper').length > 0) {
        var $container = $('.iso-box-wrapper');
        var $imgs = $('.iso-box img');

        $container.imagesLoaded(function () {
            $container.isotope({
                layoutMode: 'fitRows',
                itemSelector: '.iso-box',
            });

            $imgs.load(function () {
                $container.isotope('reLayout');
            });
        });

        $('.filter-wrapper li a').on('click', function () {
            var $this = $(this);
            var filterValue = $this.attr('data-filter');

            $container.isotope({
                filter: filterValue,
                animationOptions: {
                    duration: 750,
                    easing: 'linear',
                    queue: false,
                },
            });

            if ($this.hasClass('selected')) {
                return false;
            }

            var filterWrapper = $this.closest('.filter-wrapper');
            filterWrapper.find('.selected').removeClass('selected');
            $this.addClass('selected');
            return false;
        });
    }

    // Smooth scroll for anchors
    $('a.smoothScroll').on('click', function (event) {
        var target = this.hash;
        if (target && $(target).length) {
            event.preventDefault();
            $('html, body').animate(
                {
                    scrollTop: $(target).offset().top - 50,
                },
                700
            );
        }
    });

    // Initialize wow animations
    if (typeof WOW === 'function') {
        new WOW().init();
    }
});
