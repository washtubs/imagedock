package imagedock

import "context"

var p Providers

func setProviders(inject Providers) {
	p = inject
}

type Providers interface {
	provideTagModel(c context.Context) TagModel
}
