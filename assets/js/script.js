  function cust_includeHTML() {
      var z, i, elmnt, file, xhttp;
      /* Loop through a collection of all HTML elements: */
      z = document.getElementsByTagName("*");
      for (i = 0; i < z.length; i++) {
          elmnt = z[i];
          /*search for elements with a certain atrribute:*/
          file = elmnt.getAttribute("w3-include-html");
          if (file) {
              /* Make an HTTP request using the attribute value as the file name: */
              xhttp = new XMLHttpRequest();
              xhttp.onreadystatechange = function() {
                  if (this.readyState == 4) {
                      if (this.status == 200) {
                          elmnt.innerHTML = this.responseText;
                      }
                      if (this.status == 404) {
                          elmnt.innerHTML = "Page not found.";
                      }
                      /* Remove the attribute, and call this function once more: */
                      elmnt.removeAttribute("w3-include-html");
                      cust_includeHTML();
                  }
              }
              xhttp.open("GET", file, true);
              xhttp.send();
              /* Exit the function: */
              return;
          }
      }
  }

  function hamburger(x) {
      x.classList.toggle("change");
  }

  function MakePosNeg() {
      var TDs = document.querySelectorAll('.plusmin');

      for (var i = 0; i < TDs.length; i++) {
          var temp = TDs[i];
          if (temp.firstChild.nodeValue.indexOf('-') == 0) {
              temp.className = "text-right negative";
          } else {
              temp.className = "text-right positive";
          }
      }
  }
  onload = MakePosNeg()






  /* Set the width of the side navigation to 250px and the left margin of the page content to 250px and add a black background color to body */
  function openNav() {
      document.getElementById("mySidenav").style.width = "300px";
      //  document.getElementById("body").style.marginLeft = "250px";
      document.body.style.backgroundColor = "rgba(0,0,0,.1)";
  }

  /* Set the width of the side navigation to 0 and the left margin of the page content to 0, and the background color of body to white */
  function closeNav() {
      document.getElementById("mySidenav").style.width = "0";
      //  document.getElementById("body").style.marginLeft = "0";
      document.body.style.backgroundColor = "rgba(255,255,255)";
  }