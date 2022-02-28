package key

import (
	mock_helper "bit-driver-matching-service/domain/.mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewFileRepository(t *testing.T) {
	var fileRepo, err = NewFileRepository("/")

	assert.Nil(t, err)
	assert.NotNil(t, fileRepo)
}

func TestNewFileRepositoryFailOnEmptyFilePath(t *testing.T) {
	var fileRepo, err = NewFileRepository("")

	assert.NotNil(t, err)
	assert.Equal(t, ErrorFilePathCannotBeEmpty, err.Error())
	assert.Nil(t, fileRepo)
}

func TestFileRepository_ReadFromFile(t *testing.T) {
	var controller = gomock.NewController(t)
	defer controller.Finish()

	var expectedFileOutput = Keys{"key": "val"}

	var ioUtil = mock_helper.NewMockIOUtil(controller)
	ioUtil.EXPECT().ReadFromFile(gomock.Any()).SetArg(0, expectedFileOutput)

	var fileRepo, _ = NewFileRepository("/", ioUtil)

	var actualFileOutput, err = fileRepo.ReadFromFile()

	assert.Nil(t, err)
	assert.Equal(t, expectedFileOutput, actualFileOutput)
}

func TestFileRepository_WriteToFile(t *testing.T) {
	var controller = gomock.NewController(t)
	defer controller.Finish()

	var ioUtil = mock_helper.NewMockIOUtil(controller)
	ioUtil.EXPECT().WriteToFile(gomock.Any()).Return(nil)

	var fileRepo, _ = NewFileRepository("/", ioUtil)

	err := fileRepo.WriteToFile(Keys{"key": "val"})

	assert.Nil(t, err)
}
