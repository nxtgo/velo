# velo

a tiny and fast image proxy + transformer written in go.  
inspired by unjs/ipx, but built for simplicity, speed, and mental stability (no).

## how it works

`/q_80,gray,s_200x200?src=https://i.imgur.com/example.png`

* `q_80`: set jpeg quality to 80  
* `gray`: apply grayscale filter  
* `s_200x200`: resize to 200x200 pixels  
* `src`: the image url

## features

* built on fasthttp
* caching
* configurable domain whitelist
* simple
* stats

## env

| name | default | description |
|------|----------|-------------|
| `VELO_ADDR` | `:8080` | address to bind |
| `VELO_CACHE_DIR` | `store` | cache directory |
| `VELO_MAX_IMAGE_SIZE` | `10 << 20` | max image size in bytes |
| `VELO_WHITELISTED_DOMAINS` | `^(?:i\.imgur\.com\|cdn\.discordapp\.com)$` | regex whitelist for allowed domains |

## docker or whatever

```bash
docker build -t velo .
docker run -p 8080:8080 velo
```

# license

under CC0 1.0 universal, rights waived.
