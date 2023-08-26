# media

The package defines common vocabulary (core data types) for media objects, following the best practices from the industry.


## Resolution

The resolution defines the ontology for expressing media object variants. It is inherited from concepts used by [Digital standards](https://en.wikipedia.org/wiki/List_of_common_resolutions#Digital_standards).

It defined 9 standardizes profiles for image resolutions:
1. Half Quarter VGA (240 x 160)
2. Wide Quarter VGA (400 x 240)
3. Standard Definition (720 x 480)
4. High Definition (1280 x 720)
5. Full HD (1920 x 1080)
6. Quad HD, 2K (2560 x 1440)
7. Ultra HD, False 4K (3840 x 2160)
8. Hi-Vision System, True 4K (4096 × 2160)
9. Super Hi-Vision System, 8K (7680 × 4320)

For example, modelling of Unsplash images using the resolution types:

```go
import "github.com/fogfish/schemaext/media"

type Image struct {
  media.HQVGA    `json:"schema:thumb"`
  media.WQVGA    `json:"schema:small"`
  media.HD       `json:"schema:regular"`
  media.Origin   `json:"schema:full"`
  schemeorg.Url  `json:"schema:raw"`
}
```

Using the resolution types, the application can [standardize photo sizes](https://www.adobe.com/uk/creativecloud/photography/discover/standard-photo-sizes.html).

### Instagram

Use `media.HD` type for standard Instagram images, which are 1080 x 1080 with 1:1 aspect ratio. 

Use `media.FHD` type for Instagram stories, which are 1080 x 1920 with an aspect ratio of 9:16. 

### Facebook

Use `media.HD` type for Facebook feed posts, which are 1200 x 630.

Use `media.FHD` type for Facebook stories, which are 1080 x 1920 with an aspect ratio of 9:16. 

### Twitter

Use `media.HD` for Twitter posts including an image as part of a link. Twitter uses 1200 x 628 pixels.

Use `media.HD` for Twitter posts including single image. Twitter makes it 1200 x 675 pixels.

Use `media.SD` for Twitter posts including multiple images. Twitter makes each one 700 x 800 pixels.

### LinkedIn

Use `media.HD` for LinkedIn photos on the feed. LinkedIn uses 1200 x 627 pixels.

