package integration

import (
	"os"

	"github.com/containers/podman/v3/libpod/define"
	"github.com/containers/podman/v3/pkg/annotations"
	. "github.com/containers/podman/v3/test/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Podman container inspect", func() {
	var (
		tempdir    string
		err        error
		podmanTest *PodmanTestIntegration
	)

	BeforeEach(func() {
		tempdir, err = CreateTempDirInTempDir()
		if err != nil {
			os.Exit(1)
		}
		podmanTest = PodmanTestCreate(tempdir)
		podmanTest.Setup()
		podmanTest.SeedImages()
	})

	AfterEach(func() {
		podmanTest.Cleanup()
		f := CurrentGinkgoTestDescription()
		processTestResult(f)
	})

	It("podman inspect a container for the container manager annotation", func() {
		const testContainer = "container-inspect-test-1"
		setup := podmanTest.RunTopContainer(testContainer)
		setup.WaitWithDefaultTimeout()
		Expect(setup).Should(Exit(0))

		data := podmanTest.InspectContainer(testContainer)
		Expect(data[0].Config.Annotations[annotations.ContainerManager]).
			To(Equal(annotations.ContainerManagerLibpod))
	})

	It("podman inspect shows exposed ports", func() {
		name := "testcon"
		session := podmanTest.Podman([]string{"run", "-d", "--stop-timeout", "0", "--expose", "8080/udp", "--name", name, ALPINE, "sleep", "inf"})
		session.WaitWithDefaultTimeout()
		Expect(session).Should(Exit(0))
		data := podmanTest.InspectContainer(name)

		Expect(data).To(HaveLen(1))
		Expect(data[0].NetworkSettings.Ports).
			To(Equal(map[string][]define.InspectHostPort{"8080/udp": nil}))
	})

	It("podman inspect shows exposed ports on image", func() {
		name := "testcon"
		session := podmanTest.Podman([]string{"run", "-d", "--expose", "8080", "--name", name, nginx})
		session.WaitWithDefaultTimeout()
		Expect(session).Should(Exit(0))

		data := podmanTest.InspectContainer(name)
		Expect(data).To(HaveLen(1))
		Expect(data[0].NetworkSettings.Ports).
			To(Equal(map[string][]define.InspectHostPort{"80/tcp": nil, "8080/tcp": nil}))
	})
})
