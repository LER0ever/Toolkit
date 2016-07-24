var ajaxify = function(webpage, callback, tolerance, maxWaitTime) {
   var tolerance   = tolerance   || 6000;  // how long to wait for new request to arrive before considering page loaded
   var maxWaitTime = maxWaitTime || 15000; // how long to wait before giving up (in case of long polling, etc)

   var inprogress  = 0;
   var timeout, forcetimeout;

   forcetimeout = window.setTimeout(function(){
      window.clearTimeout(timeout);
      callback();
   }, maxWaitTime);

   var request_callback = function(){
      window.clearTimeout(timeout);
   };
   var response_callback = function(){
      timeout = window.setTimeout(function(){
         window.clearTimeout(forcetimeout);
         callback();
      }, tolerance);
   };

   var ready = function() {
      inprogress--;
      if (inprogress == 0) {
         response_callback();
      }
   }

   webpage.onResourceRequested = function (requestData, networkRequest) {
      // guard cause
      if (inprogress < 0) {
         inprogress = 0;
      }
      inprogress++;
      request_callback();
   }
   webpage.onResourceReceived = function (response) {
      // the callback may be fired multiple times (chunked response, http redirects, etc)
      if (response.stage == 'end') {
         ready();
      }
   }
   webpage.onResourceTimeout = function (response) {
      ready();
   }
   webpage.onResourceError = function (response) {
      ready();
   }
}
var webpage = require('webpage').create(),
    system  = require('system');

var url = system.args[1];
var output = system.args[2]
var size  = system.args[3] || 1280;

webpage.viewportSize = { width : size, height: size * 3/4 };
webpage.settings.userAgent = 'Mozilla/5.0';

var pagetimeout = window.setTimeout(function(){
   console.log('Unable to load ' + url);
   phantom.exit(1);
}, 60000);

//console.log("Opening " + url);
webpage.open(url, function(status){
   if (status !== 'success') {
      console.log('Unable to load ' + url);
      phantom.exit(1);
   }
   window.clearTimeout(pagetimeout);
   //console.log(url + " loaded");
   ajaxify(webpage, function(){
      //console.log('rendering to ' + output);
      webpage.render(output);
      console.log(webpage.evaluate(function () {
            var d= ''
               var c = document.querySelectorAll('div.flight-item')
               var l = c.length;
               for(var i =0;i<l;i++){
                  var airline = c[i].querySelector('div.airline-name');
                  if (/*airline.innerHTML == "东方航空"*/ 1){
                     var flyno = c[i].querySelector('span.flight-No')
                     //d = d+'China Eastern Airline '+flyno.innerHTML + '\n'
                     //if (flyno.innerHTML != "MU717") continue;
                     var seatrows = c[i].querySelectorAll('div.seat-row')
                     for (var j=0; j<seatrows.length; j++) {
                        if (parseInt( seatrows[j].querySelector('div.seat-price').querySelector('span.price').innerHTML.substring(12)) > 9500) continue;
                        d = d+flyno.innerHTML+" " +seatrows[j].querySelector('div.seat-type').innerHTML + ' '
                        d = d+ seatrows[j].querySelector('div.seat-price').querySelector('span.price').innerHTML.substring(12) + '\n' //querySelector('span.price').
                     }
                     //d=d+airline.innerHTML+'\n'
                  }
               }
            return d
         }))
      phantom.exit(0);
   });

});