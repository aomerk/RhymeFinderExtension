chrome.browserAction.onClicked.addListener(function(tab) {
  chrome.tabs.create({'url': chrome.extension.getURL('rhyme.html')}, function(tab) {
    // open rhyme finder
  });
});
