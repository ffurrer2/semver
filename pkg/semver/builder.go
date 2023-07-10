// SPDX-License-Identifier: MIT

package semver

type Builder struct {
	prefix        string
	major         uint
	minor         uint
	patch         uint
	preRelease    []string
	buildMetadata []string
}

func NewBuilder() *Builder {
	return &Builder{
		prefix:        "",
		major:         0,
		minor:         0,
		patch:         0,
		preRelease:    []string{},
		buildMetadata: []string{},
	}
}

func (b *Builder) Prefix(prefix string) *Builder {
	return &Builder{
		prefix:        prefix,
		major:         b.major,
		minor:         b.minor,
		patch:         b.patch,
		preRelease:    b.preRelease,
		buildMetadata: b.buildMetadata,
	}
}

func (b *Builder) Major(major uint) *Builder {
	return &Builder{
		prefix:        b.prefix,
		major:         major,
		minor:         b.minor,
		patch:         b.patch,
		preRelease:    b.preRelease,
		buildMetadata: b.buildMetadata,
	}
}

func (b *Builder) Minor(minor uint) *Builder {
	return &Builder{
		prefix:        b.prefix,
		major:         b.major,
		minor:         minor,
		patch:         b.patch,
		preRelease:    b.preRelease,
		buildMetadata: b.buildMetadata,
	}
}

func (b *Builder) Patch(patch uint) *Builder {
	return &Builder{
		prefix:        b.prefix,
		major:         b.major,
		minor:         b.minor,
		patch:         patch,
		preRelease:    b.preRelease,
		buildMetadata: b.buildMetadata,
	}
}

func (b *Builder) PreRelease(preRelease []string) *Builder {
	return &Builder{
		prefix:        b.prefix,
		major:         b.major,
		minor:         b.minor,
		patch:         b.patch,
		preRelease:    preRelease,
		buildMetadata: b.buildMetadata,
	}
}

func (b *Builder) PreReleaseField(preRelease string) *Builder {
	return &Builder{
		prefix:        b.prefix,
		major:         b.major,
		minor:         b.minor,
		patch:         b.patch,
		preRelease:    append(b.preRelease, preRelease),
		buildMetadata: b.buildMetadata,
	}
}

func (b *Builder) BuildMetadata(buildMetadata []string) *Builder {
	return &Builder{
		prefix:        b.prefix,
		major:         b.major,
		minor:         b.minor,
		patch:         b.patch,
		preRelease:    b.preRelease,
		buildMetadata: buildMetadata,
	}
}

func (b *Builder) BuildMetadataField(buildMetadata string) *Builder {
	return &Builder{
		prefix:        b.prefix,
		major:         b.major,
		minor:         b.minor,
		patch:         b.patch,
		preRelease:    b.preRelease,
		buildMetadata: append(b.buildMetadata, buildMetadata),
	}
}

func (b *Builder) Build() (*SemVer, bool) {
	semver := &SemVer{
		Prefix:        b.prefix,
		Major:         b.major,
		Minor:         b.minor,
		Patch:         b.patch,
		PreRelease:    b.preRelease,
		BuildMetadata: b.buildMetadata,
	}
	return semver, semver.IsValid()
}
