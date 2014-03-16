
$(function() {

  var conn;
  var traffic_count = 0;
  var msg = $("#msg");
  var log = $("#log");

  var audioContext;

  try {
    // Fix up for prefixing
    window.AudioContext = window.AudioContext||window.webkitAudioContext;
    audioContext = new AudioContext();
  }
  catch(e) {
    alert('Web Audio API is not supported in this browser');
  }

  function playAudio(data) {
    var startTime = 0;
    var fileReader = new FileReader();
    fileReader.onload = function (e) {
  
      var data = new DataView(e.target.result);

      var INCOMING_RATE = 48000;
      var OUTGOING_RATE = 48000; // 32000 is minimum
      var UPSAMPLE_RATIO = OUTGOING_RATE / INCOMING_RATE;

      var incoming_length = data.byteLength / Int16Array.BYTES_PER_ELEMENT

      var audio = new Int16Array(incoming_length * UPSAMPLE_RATIO);
      var len = audio.length;

      var gg = 0
      for (var jj = 0; jj < incoming_length; jj++) {
          // TODO -- to get rid of the aliasing effect this has on the audio
          // some sort of interpolation between the samples should be applied here.
          for (var hh = gg; hh < gg+UPSAMPLE_RATIO; hh++) {
              audio[hh] = data.getInt16(jj * Int16Array.BYTES_PER_ELEMENT, false);
          }   
          gg += UPSAMPLE_RATIO
      }

      var right = new Float32Array(audio.length);
      var left = new Float32Array(audio.length);

      var channelCounter = 0;
      for (var i = 0; i < audio.length-1; i += 2) {
          var normalizedRight = audio[i] / 32768
          var normalizedLeft  = audio[i+1] / 32768
          right[channelCounter] = normalizedRight;
          left[channelCounter] = normalizedLeft
          channelCounter++;
      }

      var source = audioContext.createBufferSource();

      var audioBuffer = audioContext.createBuffer(2, right.length, OUTGOING_RATE);
      audioBuffer.getChannelData(0).set(right);
      audioBuffer.getChannelData(1).set(left);

      source.onended = function() {
        $("#tracker button").text("▶");
      }

      source.buffer = audioBuffer;

      source.connect(audioContext.destination);

      source.start(startTime);
      startTime += audioBuffer.duration;

      // todo -- idea is to use braile unicode as progress spinner indicator for network activity
      // ⣾⣽⣻⢿⡿⣟⣯⣷

    };
    fileReader.readAsArrayBuffer(data)
  }

  playTheseLicks = function(licks) {
    if (!conn || !msg.val()) {
      alert('no conn or no msg val');
    }
    $("button").text("❚ ❚")
    sendThroughSocket(JSON.stringify(licks));
  }

  playTestSequence = function() {
    var licks = new Array();

    note1 = {}
    note1["tick"] = 0
    note1["note"] = "c4"
    note1["dura"] = 10
    note1["modd"] = 100
    note1["velo"] = 100

    note2 = {}
    note2["tick"] = 2
    note2["note"] = "e4"
    note2["dura"] = 10
    note2["modd"] = 100
    note2["velo"] = 100

    note3 = {}
    note3["tick"] = 4
    note3["note"] = "g4"
    note3["dura"] = 10
    note3["modd"] = 100
    note3["velo"] = 100

    licks = [note1, note2, note3]

    playTheseLicks(licks)
  }

  sendThroughSocket = function(data) {
    moreTraffic();
    conn.send(data);
  }

  moreTraffic = function() {
    traffic_count++;
    if(traffic_count >= 0) {
      $("#network-indicator").show();
    }
  }

  lessTraffic = function() {
    traffic_count--;
    if(traffic_count <= 0) {
      $("#network-indicator").hide();
    }
  }

  if (window["WebSocket"]) {
      conn = new WebSocket("ws://localhost:8080/ws");
      conn.onclose = function(evt) {
        console.log('closed')
      }
      conn.onmessage = function(evt) {
        lessTraffic();
        playAudio(evt.data)
      }
  } else {
      alert($("<div><b>Your browser does not support WebSockets.</b></div>"))
  }
});
