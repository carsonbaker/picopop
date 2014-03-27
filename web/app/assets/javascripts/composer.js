$(function() {

  var licks = {}
  var note_licks = {}

  var cell_width = 20; // this is the width of a 16th note
  var cell_height = 12; 
  var half_cell_height = cell_height / 2;
  var half_cell_width = cell_width / 2;

  var selected_cell_color = "#fee";
  var normal_cell_color = "#49a";

  var current_selection;
  var is_dragging = false;

  var mouse_x, mouse_y;

  var view_height = $("#sequencer").height();
  var view_width = $("#sequencer").width();


  $("#tracker").submit(function() {
    play_all_notes();
    event.preventDefault();
  });

  var note_identifiers = ["b5", "a5s", "a5", "g5s", "g5", "f5s", "f5", "e5", "d5s", "d5", "c5s", "c5", "b4", "a4s", "a4", "g4s", "g4", "f4s", "f4", "e4", "d4s", "d4", "c4s", "c4", "b3", "a3s", "a3", "g3s", "g3", "f3s", "f3", "e3", "d3s", "d3", "c3s", "c3", "b2", "a2s", "a2", "g2s", "g2", "f2s", "f2", "e2", "d2s", "d2", "c2s", "c2", "b1", "a1s", "a1", "g1s", "g1", "f1s", "f1", "e1", "d1s", "d1", "c1s", "c1"]

  // initialize an inversion of the note map, for convenience
  var note_inverse_map = {}
  var note_inverse_i = 0
  for(note in note_identifiers) {
    note_inverse_map[note_identifiers[note]] = note_inverse_i++
  }

  // initialize the note_licks hash to have an empty hash for each note
  note_identifiers.forEach(function(note) {
    note_licks[note] = {}
  });

  play_all_notes = function() {
    var lick_array = Object.keys(licks).map(function (key) { return licks[key]; });
    playTheseLicks(lick_array);
  }

  half_step_up = function(note) {
    return note_identifiers[note_inverse_map[note] - 1] || note_identifiers[0]
  }

  half_step_down = function(note) {
    return note_identifiers[note_inverse_map[note] + 1] || note_identifiers[note_identifiers.length-1]
  }

  t = $("#time-indicators")[0].getContext("2d")
  c = $('#composer')[0].getContext("2d");

  drawLines = function(c) {

    c.lineWidth = 1.0;

    for (var x = 0; x < view_width; x += cell_width) {

      if(x % 320 == 0) {
        stroke_style = '#b9e3ff';
      }
      else if(x % 80 == 0) {
        stroke_style = '#555';
      } else {
        stroke_style = '#333';
      }

      pos_x = x + 0.5;

      for (var y = 0; y < (12 * 5) + 1; y++) {

        if(y % 12 == 0) {
          c.strokeStyle = '#fff';
        } else {
          c.strokeStyle = stroke_style;
        }

        pos_y = y * cell_height + 0.5;
        c.beginPath();
        c.moveTo(pos_x, pos_y);
        c.lineTo(pos_x, pos_y+0.5);
        c.stroke();
        c.closePath();
      }
    }
  }

  var background_grid_buffer = renderToCanvas(view_width, view_height, drawLines);

  drawLicks = function() {
    for(var key in licks) {
      lick = licks[key]
      if(lick["sele"]) {
        c.fillStyle = selected_cell_color
      } else {
        c.fillStyle = normal_cell_color
      }
      pos = pos_for_lick(lick)
      c.beginPath();
      c.rect(pos[0] + 0.5, pos[1] + 0.5, cell_width * lick["dura"], cell_height);
      c.closePath();
      // roundRect(c, pos[0] + 0.5, pos[1] + 0.5, cell_height * lick["dura"], cell_height, 10)
      c.stroke();
      c.fill();
    }
  }

  drawTimeMarkers = function() {
    t.fillStyle = "#fff"
    t.font = "11px Arial";
    for (var x = 0; x < view_width; x ++) {
      t.fillText(x, x * cell_width * 16 + 18, 14);
    }
  }

  redraw = function(x) {
    c.clearRect(0, 0, view_width, view_height);
    c.drawImage(background_grid_buffer, 0, 0);
    drawLicks();
    drawTimeMarkers();
  }

  pos_for_lick = function(lick) {
    return [lick['tick'] * cell_width, note_identifiers.indexOf(lick['note']) * cell_height]
  }

  note_for_pos = function(y) {
    return note_identifiers[Math.floor(y / cell_height)]
  }

  tick_for_pos = function(x) {
    return Math.floor(x / cell_width);
  }

  move_lick = function(lick, new_tick, new_note) {

    // start by deleting the old lick from both tracking hashes
    delete_lick(lick)

    // update the attributes
    lick["tick"] = new_tick
    lick["note"] = new_note

    // now add the lick back, with attributes updated to the new spot.
    licks[new_tick + new_note] = lick
    note_licks[new_note][new_tick] = lick

  }

  delete_lick = function(lick) {
    // remove from both tracking hashes
    delete licks[lick["tick"] + lick["note"]]
    delete note_licks[lick["note"]][lick["tick"]] 

  }

  update_note_on_drag = function(e) {
    // document.body.style.cursor = 'move';
    drag_tick = tick_for_pos(e.offsetX)
    drag_note = note_for_pos(e.offsetY)

    // if we haven't moved the lick to another part of the grid, just stop
    if(drag_tick == current_selection["tick"] && drag_note == current_selection["note"]) { return }

    is_dragging = true;

    // otherwise let's update it, and redraw it

    if(licks[drag_tick+drag_note]) {
      // there's already a note on the spot that we're dragging over
      // current way of dealing with this is to just prevent
      // the human from sequencing a note twice in the same spot.
      return
    } else {
      move_lick(current_selection, drag_tick, drag_note)
    }

    redraw();
  }

  $("body").on("keydown", function(e) {

    var x = mouse_x
    var y = mouse_y

    switch(e.keyCode) {
      case 49: // 1 (add whole note)
        lick = add_select_and_play_lick(x, y, 16)
        break;
      case 50: // 2 (add half note)
        lick = add_select_and_play_lick(x, y, 8)
        break;
      case 51: // 3 (add quarter note)
        lick = add_select_and_play_lick(x, y, 4)
        break;
      case 52: // 4 (add eighth note)
        lick = add_select_and_play_lick(x, y, 2)
        break;
      case 53: // 4 (add sixteenth note)
        lick = add_select_and_play_lick(x, y, 1)
        break;
      case 32: // play (space)
        play_all_notes();
        event.preventDefault();
        break;
      case 88: // delete (x)
        delete_lick(current_selection)
        break;
      case 76: // right (l)
        move_lick(current_selection, current_selection["tick"] + 1, current_selection["note"])
        break;
      case 72: // left (h)
        move_lick(current_selection, current_selection["tick"] - 1, current_selection["note"])
        break;
      case 75: // up (k)
        move_lick(current_selection, current_selection["tick"], half_step_up(current_selection["note"]))
        break;
      case 74: // down (j)
        move_lick(current_selection, current_selection["tick"], half_step_down(current_selection["note"]))
        
        break;
    }

    redraw();
  });

  $("#composer").on("mouseout", function(e) {
    document.body.style.cursor = 'default';
    is_dragging = false;
    this.onmousemove = null;
  });

  $("#composer").on("mouseover", function(e) {
    document.body.style.cursor = 'crosshair';
  });

  $("#composer").on("mousemove", function(e) {
    mouse_x = e.offsetX;
    mouse_y = e.offsetY;
  })

  $("#composer").on("mouseup", function(e) {
    this.onmousemove = null;
    quick_lick = {}
    $.extend(quick_lick, current_selection, {"tick":0});
    if(is_dragging) {
      is_dragging = false;
      playTheseLicks([quick_lick]);
    }
    document.body.style.cursor = 'crosshair';
  });

  lick_hit_detection = function(x, y) {

    click_tick = tick_for_pos(x)
    click_note = note_for_pos(y)

    for(tick_key in note_licks[click_note]) {
      if(tick_key <= click_tick) {
        lick_from_note = note_licks[click_note][tick_key]
        if(click_tick <= (lick_from_note["tick"] + lick_from_note["dura"])) {
          return lick_from_note
        }
      }
    }

  }

  add_a_lick = function(x, y, duration) {

    click_tick = tick_for_pos(x)
    click_note = note_for_pos(y)

    lick = {}
    lick["tick"] = click_tick
    lick["note"] = click_note
    lick["dura"] = duration
    lick["modd"] = 100
    lick["velo"] = 100
    lick["sele"] = true
    licks[click_tick+click_note] = lick
    note_licks[click_note][click_tick] = lick

    return lick
  }

  add_select_and_play_lick = function(x, y, duration) {

    // deselect whatever is selected
    if(current_selection) {
      current_selection["sele"] = false
    }

    lick = add_a_lick(x, y, duration)

    // not a selection. we're adding a the lick to the sequence
    current_selection = lick

    // quickly play a test lick so that the human can hear what he's doing
    quick_lick = {}
    $.extend(quick_lick, lick, {"tick":0});
    playTheseLicks([quick_lick]);
  }

  $("#composer").on("mousedown", function(e) {

    // is this a selection?
    lick = lick_hit_detection(e.offsetX, e.offsetY)

    if(lick) {
      // yep, cool we're making a selection now

      // we can safely deselect what was selected
      // because now it's as if we're reselecting it.
      if(current_selection) {
        current_selection["sele"] = false
      }

      lick["sele"] = true
      current_selection = lick

    } else {
      // not a selection. add a quarter note!
      add_select_and_play_lick(e.offsetX, e.offsetY, 4)

    }

    // let's add a handler to onmousemove so that we can drag the new note or selection
    this.onmousemove = update_note_on_drag

    redraw();

  });

  // this is to draw the screen initially
  // it is only called once
  redraw();

});
