# Archive Resource

Simply downloads and extracts an archive to the destination.

## Source Configuration

* `uri`: *Required.* The location of the file to download.

## Behavior

### `check`: Not implemented.

As this resource is mainly used for one-off downloads (with
[Fly](https://github.com/concourse/fly)), there aren't really any versioning
semantics.


### `in`: Download and extract the archive.

Fetches a `.tar.gz` file from the URL, and extracts it to the destination as
it's downloading.


#### Parameters

*None.*


### `out`: Not implemented.

Currently there is no output functionality. In principle, this could be
configured with a directory to compress and upload to the `uri`, however
this is not currently implemented.

#### Parameters

*None.*
