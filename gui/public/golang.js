/* eslint-disable no-unused-vars */
/*
From js to go 
- Clicked Download
- Changed a settings?
- 

From go to js
- DownloadedMetadata
- DownloadProgress
- 

*/

// object with icon url link and array of versions 
var DownloadMetadata;
var DownloadMetadataEvent = new Event("DownloadMetadataEvent");

var DownloadProgress = 0;
var DownloadProgressEvent = new Event("DownloadProgressEvent");

var DownloadedVersions = {};