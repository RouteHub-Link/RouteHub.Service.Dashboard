function initializeShadowRoot(id) {
  var shadowHost = document.getElementById(id);
  var cloneInside = shadowHost.innerHTML;
  shadowHost.innerHTML = "";

  // Create shadow root
  var shadow = shadowHost.attachShadow({ mode: "open" });

  // Add initial content
  shadow.innerHTML = `
     <link rel="stylesheet" href="/static/css/bulma.host.css" />
     <div id="content-container" style="color:var(--bulma-body-color) !important;">
       ${cloneInside}
     </div>
   `;

  // Function to initialize feather icons for this specific shadow root
  function initializeFeatherForShadow() {
    try {
      // Force feather to work inside shadow DOM
      var icons = Array.from(shadow.querySelectorAll("[data-feather]"));

      icons.forEach(function (element) {
        var iconName = element.getAttribute("data-feather");
        var svgString = feather.icons[iconName].toSvg();
        var tempDiv = document.createElement("div");
        tempDiv.innerHTML = svgString;
        var svgElement = tempDiv.firstChild;

        // Copy over any additional attributes from original element
        Array.from(element.attributes).forEach(function (attr) {
          if (attr.name !== "data-feather") {
            svgElement.setAttribute(attr.name, attr.value);
          }
        });

        element.parentNode.replaceChild(svgElement, element);
      });

      //console.log( "Feather icons initialized for shadow root:", id, "Found icons:", icons.length, );
    } catch (error) {
      console.error("Error initializing Feather icons:", error);
    }
  }

  // Check if Feather is already loaded
  if (window.feather) {
    setTimeout(initializeFeatherForShadow, 100);
  } else {
    // Load Feather only if it's not already loading or loaded
    var existingScript = document.querySelector('script[src*="feather"]');
    if (!existingScript) {
      var featherScript = document.createElement("script");
      featherScript.src =
        "https://cdn.jsdelivr.net/npm/feather-icons/dist/feather.min.js";
      featherScript.defer = true;
      document.body.appendChild(featherScript);

      featherScript.onload = function () {
        setTimeout(initializeFeatherForShadow, 100);
      };
    } else {
      existingScript.addEventListener("load", function () {
        setTimeout(initializeFeatherForShadow, 100);
      });
    }
  }

  return shadow;
}
