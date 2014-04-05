import 'package:polymer/polymer.dart';
import 'dart:html';
import 'dart:typed_data';

/**
 * A Polymer click counter element.
 */
@CustomTag('ws-counter')
class Counter extends PolymerElement {
  @observable String x = 'No data received yet';
  WebSocket ws;
  int cnt = 0;
  Counter.created() : super.created() {
    ws = new WebSocket('ws://localhost:1234/tabledata');
    ws.onMessage.listen((MessageEvent e) {
      var reader = new FileReader();
      reader.onLoadEnd.listen((_) { 
        var buffer = reader.result as ByteBuffer;
        var lst = new Uint64List.view(buffer);
        x = "Received data : ${lst.toString()}, len = ${lst.length}, lenbytes = ${lst.lengthInBytes}";
        
      });
      reader.readAsArrayBuffer(e.data);
    } );
    ws.onOpen.listen((e) {
      x = 'Connected';
    });

    ws.onClose.listen((e) {
      
      x = 'Closed';
    });

    ws.onError.listen((e) {
      x = 'Error';
    });
  }

}

