	"github.com/stretchr/testify/mock"
	"github.com/golangci/golangci-lint/pkg/result"
	log := logutils.NewMockLog()
	log.On("Infof", "The diff contains only additions: no original or deleted lines: %#v", mock.Anything)
