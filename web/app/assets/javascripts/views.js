
// keep the time indicators on the top of the viewport always
$(window).scroll(function(){
  $('#time-indicators').css({
      'top': Math.max(0, $(this).scrollTop())
  });
});
