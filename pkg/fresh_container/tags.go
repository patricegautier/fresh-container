package fresh_container

import (
	"github.com/blang/semver"
	log "github.com/sirupsen/logrus"
	"strings"
)

//func ImageTags(ctx context.Context, cfg *config.Config, image Image) ([]string, error) {
//  // Create the registry client.
//  r, err := createRegistryClient(ctx, image.Domain, cfg)
//  if err != nil {
//    return []string{}, err
//  }

//  tags, err := r.Tags(ctx, image.Path)
//  if err != nil {
//    return []string{}, err
//  }
//  sort.Strings(tags)

//  return tags, nil
//}

func TagsToVersions(tags []string, tagPrefix string, skipInvalid bool, prefixAlreadyStripped bool) (versions semver.Versions, err error) {
	for _, tag := range tags {
		if prefixAlreadyStripped || tagPrefix == "" || strings.HasPrefix(tag, tagPrefix) {
			// only consider tags that have the specified prefix
			tag = strings.TrimPrefix(tag, tagPrefix)
			v, err := semver.Parse(tag)
			if err == nil {
				versions = append(versions, v)
			} else if skipInvalid {
				log.WithFields(log.Fields{
					"tag":       tag,
					"tagPrefix": tagPrefix,
					"error":     err}).Warn("Skipping image tag")
			} else {
				return semver.Versions{}, err
			}
		}
	}
	return versions, nil
}

//func ImageTag(img string) (semver.Version, error) {
//  image, err := registry.ParseImage(img)
//  if err != nil {
//    return semver.Version{}, err
//  }

//  return semver.Parse(image.Tag)
//}
