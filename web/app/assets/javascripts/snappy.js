// $(function() {

//   // First lets create our drawing surface out of existing SVG element
//   // If you want to create new surface just provide dimensions
//   // like s = Snap(800, 600);
//   var s = Snap("#composer");
//   // Lets create big circle in the middle:

//   var piano_x_inc = 40;

//   var white_keys = s.group();

//   for (var key_i = 0; key_i < 12 * 4; key_i++) {
//     pos = key_i * piano_x_inc;
//     white_keys.add(s.line(pos, 0, pos, 800));
//   }

//   // var black_keys = s.group(keys[1],keys[3],keys[6],keys[8],keys[10]);

//   white_keys.attr({
//       fill: "#f00"
//   });

//   // black_keys.attr({
//   //   fill: "l(0,0,0,1)#fff-#ddd",
//   // });
    
// });
