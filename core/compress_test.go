package core

import (
	"testing"
)

var testInput = []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam ullamcorper tempus sagittis. Vestibulum rutrum nisi finibus mollis rhoncus. Aliquam erat volutpat. Sed porttitor ex eget elementum lobortis. Nullam ut dolor at sapien vestibulum fermentum ac eget ante. Mauris convallis dui eu laoreet bibendum. Mauris ante arcu, porta et lacus id, cursus sodales justo. Mauris sed nisi vehicula, lacinia ligula mattis, semper arcu. Pellentesque in molestie diam, non pretium lacus. Integer ante augue, porttitor a lobortis in, pharetra nec diam. Sed scelerisque purus a neque faucibus, sit amet commodo nisi tincidunt. Etiam sapien nibh, venenatis quis convallis sed, scelerisque sit amet ante. Integer urna odio, suscipit sit amet sapien nec, convallis consectetur massa. Aliquam malesuada lectus justo, vitae sollicitudin mauris mattis sit amet. Mauris condimentum iaculis interdum. Aliquam iaculis leo mauris, sit amet dictum ex ultricies in. Aliquam a enim vel neque porta eleifend ut ac quam. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec commodo massa et varius ultrices. Duis varius lobortis ex, ac eleifend odio venenatis ut. Phasellus non luctus eros. Sed sodales lectus id odio porta porta. Curabitur id lacinia quam. Etiam finibus nisi quis velit dapibus auctor. Integer sollicitudin, felis et ornare iaculis, diam risus hendrerit lacus, vel lacinia sem dui vitae erat.")

func TestCompressDecompressBytes(t *testing.T) {
	compressed := CompressBytes(testInput)
	t.Log("compressed: " + string(compressed) + "\n\n")

	decompressed := DecompressBytes(compressed)
	t.Log("decompressed: " + string(decompressed) + "\n\n")
}
