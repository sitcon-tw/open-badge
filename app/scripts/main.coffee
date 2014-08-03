window.onload = ->
  connectButton = document.getElementById("connect")
  connectButton.addEventListener 'click', (e) ->
    OpenBadges.connect({
        callback: "http://badge.sitcon.com/callback",
        scope: ['issue']
    });
    e.preventDefault()
