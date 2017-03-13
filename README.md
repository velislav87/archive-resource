# Archive Resource

# Forked to introcude versioning functionality based on the url endpoint
# Similar to the FTP resource type

Simply downloads and extracts an archive to the destination.

**NOTE**: This resource is only intended for use in `fly execute` (it's how
your inputs get uploaded). It won't work in a pipeline because `check` never
yields any valid versions. This is because a download URL is not enough to
continuously integrate with something, since the endpoint isn't versioned.
You probably want the [S3 resource](https://github.com/concourse/s3-resource)
or the [GitHub Release
resource](https://github.com/concourse/github-release-resource) instead.

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
