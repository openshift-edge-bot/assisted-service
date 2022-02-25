package ignition

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/host/hostutil"
	"github.com/openshift/assisted-service/models"
	"github.com/openshift/assisted-service/pkg/s3wrapper"
	"github.com/sirupsen/logrus"
)

type dummyGenerator struct {
	log      logrus.FieldLogger
	workDir  string
	cluster  *common.Cluster
	s3Client s3wrapper.API
}

// NewDummyGenerator returns a Generator that creates the expected files but with nonsense content
func NewDummyGenerator(workDir string, cluster *common.Cluster, s3Client s3wrapper.API, log logrus.FieldLogger) Generator {
	return &dummyGenerator{
		workDir:  workDir,
		log:      log,
		cluster:  cluster,
		s3Client: s3Client,
	}
}

// Generate creates the expected ignition and related files but with nonsense content
func (g *dummyGenerator) Generate(_ context.Context, installConfig []byte, platformType models.PlatformType) error {
	installConfigPath := filepath.Join(g.workDir, "install-config.yaml")
	err := ioutil.WriteFile(installConfigPath, installConfig, 0600)
	if err != nil {
		return err
	}

	toUpload := fileNames[:]
	for _, host := range g.cluster.Hosts {
		toUpload = append(toUpload, hostutil.IgnitionFileName(host))
	}

	for _, fileName := range toUpload {
		f, err := os.Create(filepath.Join(g.workDir, fileName))
		if err != nil {
			return err
		}
		defer f.Close()
		data := "data"
		// use the pre-baked data
		if fileName == "kubeconfig-noingress" {
			data = kubeconfig
		} else if fileName == "bootstrap.ign" {
			data = bootstrapIgnition
		}
		_, err = f.WriteString(data)
		if err != nil {
			return err
		}
	}
	return nil
}

// UploadToS3 uploads the generated files to the configured S3-compatible storage
func (g *dummyGenerator) UploadToS3(ctx context.Context) error {
	return uploadToS3(ctx, g.workDir, g.cluster, g.s3Client, g.log)
}

func (g *dummyGenerator) UpdateEtcHosts(serviceIPs string) error {
	return nil
}

const kubeconfig string = `
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURRRENDQWlpZ0F3SUJBZ0lJZGd5RFlHbkRoVkl3RFFZSktvWklodmNOQVFFTEJRQXdQakVTTUJBR0ExVUUKQ3hNSmIzQmxibk5vYVdaME1TZ3dKZ1lEVlFRREV4OXJkV0psTFdGd2FYTmxjblpsY2kxc2IyTmhiR2h2YzNRdApjMmxuYm1WeU1CNFhEVEl3TURVeU5qRTFNRFV5TjFvWERUTXdNRFV5TkRFMU1EVXlOMW93UGpFU01CQUdBMVVFCkN4TUpiM0JsYm5Ob2FXWjBNU2d3SmdZRFZRUURFeDlyZFdKbExXRndhWE5sY25abGNpMXNiMk5oYkdodmMzUXQKYzJsbmJtVnlNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQTNzdnNGbVlYd3pNMApIcDNDY1p3Z2RBSnJWMXRocDNxMzcyMFV1RkdIUitvSXhhNkkwQXBVcnlxalhnY3JBSXRSTzBOZlkzMGUvaXdICmZ6bXI0MTZ4YVg0S3JXNUl1SkIyZkQ1OVFUQVBPRUJJSWhrRkV2Z0cxMitQaGlOVmhsbGt2V0oyTGVIL1V3cjMKN1J3eHlERFFrQ3ZDR1o5Ky83djlmSmVVVnJhSW5sL1dYTGp0c2orRXowRm1Ucks3bXpKaVFQcGs1RjNmUHpyMwpHQlh6REZQSU9JeHR6TTdsSng2dFNlaG5hZlByTHMxb3lXTm1HdXRNK0trYzE1dUROWUhybUo4Q1ZqZm83a2Q0CnQ3SXR6NUN5TGxwRFdwVEdEdmVoOXdMRnpRVjN2NHA0amJNZlZVTVJiV1BFRHp4cWphVGIwalVvVUFBRlBicTkKUTMxNG1zWlF3d0lEQVFBQm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FxUXdEd1lEVlIwVEFRSC9CQVV3QXdFQgovekFkQmdOVkhRNEVGZ1FVWkNGdmF3QWdWSnZtcGhIT20vbEdlZXdjZzI4d0RRWUpLb1pJaHZjTkFRRUxCUUFECmdnRUJBTmNxY0ZaMEJnZmM0TnlhclMwNXRWeGkzbzU2Qi9KMmpwNnVRVUtvZlJyMGhraUhYVDd0QnZubEppUSsKeSt5OHBEb29pVHpJOHMrMFd1dEN3S1d5dHdXNnNIczd4Zk9MUEtncFZXZWgyekRyZDhHY29hbXhlMzRmalRwdgppeUxDeWN2TkIzb1FpWkcyWkU1SDRXTmMwUmZXU0ZHUDcweG1JKzVVY2RvZTR6R3lEQU43bm5CYmZtU2xIK0kwCkkzWC9FaS9lRnZka081eG9kYVlUbFNWV2hTOHdib3FDVjF0SEN0YmIyQjYvY1VWTTRXUHg2Nm81Q05PMmYzMHoKcWJMU3F4ZzhYUXhjYW5FTGsxaGhNTjVlZDZkanc2YVFZUWxGSXBlcklPVVVTNDNJMGpRZTBYdmsrWWJhZGJscgpUVVR2VVV6RTdTNW5MUGRKK0hYZ2hZdFBHR1k9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0KLS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURURENDQWpTZ0F3SUJBZ0lJSXFpUUtKdUhNY293RFFZSktvWklodmNOQVFFTEJRQXdSREVTTUJBR0ExVUUKQ3hNSmIzQmxibk5vYVdaME1TNHdMQVlEVlFRREV5VnJkV0psTFdGd2FYTmxjblpsY2kxelpYSjJhV05sTFc1bApkSGR2Y21zdGMybG5ibVZ5TUI0WERUSXdNRFV5TmpFMU1EVXlOMW9YRFRNd01EVXlOREUxTURVeU4xb3dSREVTCk1CQUdBMVVFQ3hNSmIzQmxibk5vYVdaME1TNHdMQVlEVlFRREV5VnJkV0psTFdGd2FYTmxjblpsY2kxelpYSjIKYVdObExXNWxkSGR2Y21zdGMybG5ibVZ5TUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQwpBUUVBM3FDYjEycHFMcEt3M1dTVU1MclhTdFcza2ltYllVWGk1R1AyeXJhZjFYOE5pMlZwZnZHcUxtaWFpNFhVCjF2WkJhdEhYSE1SS1VteDczUTA3ZWVKUnQzNTdSUU00ZW9tbVhqQUloQjhPRU1JV0gzTTN1dTFnVlUvZkpNb2wKTlozY3JvN1BjUGFtVmRzdzIwYjhSTEIwcW55NW9FSWdLU2F5YkdGWlF6V2t4WitaR2NlNWxLbmZ0OGZpci9Fegp3NkpOK0lxbW1YOTFtVXR6T25CaFZNNFB5RzFHbnUyZ1g1SXNUOWFHbGdwOWlTSENJMnZyMEtORXZVZFJaY0k2Ck9SM0ZQUSs1WnQ5c2hPOHhYZlJpMFF3SmN2bG5xdCtDS2N6RFcwelUwc1BhVzJ5VlJzUmIyQjZFVHExMGxpQXoKN3NKRzdFc3U4TW5JVnQ1dFpqMmoxSkpUMVFJREFRQUJvMEl3UURBT0JnTlZIUThCQWY4RUJBTUNBcVF3RHdZRApWUjBUQVFIL0JBVXdBd0VCL3pBZEJnTlZIUTRFRmdRVXZ4RGJ5NnJaNWRkczBTOEtDR3dNcVRhWitFb3dEUVlKCktvWklodmNOQVFFTEJRQURnZ0VCQUtIdmxZR2N2NTFyQW1QalR2dWFIZUpaQjhoRmlsTkdHdnIxK1QxajZIY2YKcndPYzYxSCtGUjJtenNqRFBMSWhBSjV1WWlidUZEQVB3eldDUXpJT0lxeDZhanhaZlhaSTNBcDdibDB1TlBtUwpEeW1XdnJVemx2MGZoMHVwUHE1RGFiNXFPMnRPTWJxTjNjVDQxNXU0YldWaThnVmZ2Mk9SWFUxVEhpWkliVUVJCkVNbE1vRStSRU5WQzZlU0U3Rmd3S3lUQ2M2cmZ4ZGtyMU8yOEhYOFNOOGhydW9WYXdJajN6Z1k5ZWVlZUpxdE0KSmdPVVRwWThKSWIzRzcwZjRId01xQ3hrcERqUWROVTVMYTR1OUlBN0NlQldiemdrTVBiMmI4bmRMUngzQ0Z2MAovcmF1U21GY2l0MnVid1BYbWIzalBFcVhuWXh1ckp0YWtvVTAvQnRraGdrPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLQpNSUlETWpDQ0FocWdBd0lCQWdJSUhTRXB2RFlvcHBnd0RRWUpLb1pJaHZjTkFRRUxCUUF3TnpFU01CQUdBMVVFCkN4TUpiM0JsYm5Ob2FXWjBNU0V3SHdZRFZRUURFeGhyZFdKbExXRndhWE5sY25abGNpMXNZaTF6YVdkdVpYSXcKSGhjTk1qQXdOVEkyTVRVd05USTNXaGNOTXpBd05USTBNVFV3TlRJM1dqQTNNUkl3RUFZRFZRUUxFd2x2Y0dWdQpjMmhwWm5ReElUQWZCZ05WQkFNVEdHdDFZbVV0WVhCcGMyVnlkbVZ5TFd4aUxYTnBaMjVsY2pDQ0FTSXdEUVlKCktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQU1nREt3Ni9ONkY2UXBWZEJRdXpHSy95RFhkUGhtRXgKNWNiK2h6ZEw2RDFPNWZVRHBjZitHT1FEdjB2dUZrNXcveDlVRmxoRldRTkNHOEtQbWxvSFhkMHIxQllWYWc0VgpYcGFYYklGUHJmMzVGeFhzZk1HS0JUQllaRXNiKzJRTWFNci9Ec0g0cTQ5Z294NDJ5U24rV2prbnNCWnZ4Ukd0CndBT2t0b1VWUVloaWxrWDg5dHl1VjJlTU1CbmVEZmdyazZ1WEt4YjIzTWg2WEZvNHAvMW9MV1VrVmxQQ3hlM2MKNWI2VWNON29NdTI3M016enJnZzVDOXdmbGVycWFoamxBRXFHY0JMQ3NrT0NBbGhMWGRvMW5FZ09lM0ZMUHJFeApiNW93UTJkNnQvbEFlellhbU5lTUl4bTN1OWx6djJXVytLUm4rRmcxci9UZWt4Qi9ZUUVkNHhjQ0F3RUFBYU5DCk1FQXdEZ1lEVlIwUEFRSC9CQVFEQWdLa01BOEdBMVVkRXdFQi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZHN0QKOXc2UEloYUxEYTRIaU05MVNMdUw3bGlUTUEwR0NTcUdTSWIzRFFFQkN3VUFBNElCQVFBV2JZeENVUzF2SStEUQpvK1JJZmxRSU5KZFcwZkF5WEFtVUdPZGdiVW8zL2R5UldnY0RDcFdCbVhMWlhkaVJwd3R5MWV2NTNzM1VrM0ZrCldoRGJSZVFPb29nSW1RMmVZdnZxWjFlbHhGeHdYdHZtVGw1cGZ1aFdMSDI4UC8rZnZFQW04T1dpTmtSUmpsNWYKMkhLUEU1MzBMQmtsSGx0OE9naUh0bS9XOGRUcTk3Y3lXaVlsdlREeFRkREFPRTVLZERrNXpxbXRBWUpWRlhnWApyajR0UUthTkphRWdyK2JkamhKbUIyUjU4b3RxMExXeS9uNnBjTWNmSFk2Q2k5dWhwT0Y3R3RWYjdTcUNGd0tPClMvdlVORFBsK2QyRjdLMDQweUVRUExQZHI5cEF5QlNWSkZyL2lnRkJycHJJVEdkYUZ1dFhQUHYybjBtTHBvS0kKVTlIV2FJcTgKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    server: https://api.test-infra-cluster.redhat:6443
  name: test-infra-cluster
contexts:
- context:
    cluster: test-infra-cluster
    user: admin
  name: admin
current-context: admin
preferences: {}
users:
- name: admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURaekNDQWsrZ0F3SUJBZ0lJUnljNWdMZE1sMjB3RFFZSktvWklodmNOQVFFTEJRQXdOakVTTUJBR0ExVUUKQ3hNSmIzQmxibk5vYVdaME1TQXdIZ1lEVlFRREV4ZGhaRzFwYmkxcmRXSmxZMjl1Wm1sbkxYTnBaMjVsY2pBZQpGdzB5TURBMU1qWXhOVEExTWpaYUZ3MHpNREExTWpReE5UQTFNamRhTURBeEZ6QVZCZ05WQkFvVERuTjVjM1JsCmJUcHRZWE4wWlhKek1SVXdFd1lEVlFRREV3eHplWE4wWlcwNllXUnRhVzR3Z2dFaU1BMEdDU3FHU0liM0RRRUIKQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUURXR0h2ak13Tzh6SGRPTGtJcmVYMFYvdXptMjZ6dlJYQWJ2K3R6YU42dgpXbFByTFZUSkRwVHJwUVpmbnI4ZlBXQ0xSazRmRlNTZlZlR2EvRi80MzNUNDdTWGpQcHhYWEV3Y1VhYi8yRjhhCnlmdlphVGRXRHBNd3NFNzYzeXBEODhvalNQQVlRdDRqSGhRcEV3TExqVVR6VXdkTGJjdnREVEQyWmtwTE1oZFcKMmhrNGU1eVZXSzhrUDRsSUxYbnJMbFJxR2VUNGRtbmN6NHdMcTZSWXgrR0hUaFhNUzlTZEhkZ2V1dHlVZGlVMQpPNFRIOGNEZklrRnQ5Y3lkSmNZajdmREh6WWRQemlVdmtrcFJ0eHNuamlFTzVwM0h0cDAyL3kvUmlpN2t5TjByCndlNGM2Y29LZzRXZzYrUzVzaEpZV3d0ZGsrTmkxYU1BTnVIOExMbGhaUnlEQWdNQkFBR2pmekI5TUE0R0ExVWQKRHdFQi93UUVBd0lGb0RBZEJnTlZIU1VFRmpBVUJnZ3JCZ0VGQlFjREFRWUlLd1lCQlFVSEF3SXdEQVlEVlIwVApBUUgvQkFJd0FEQWRCZ05WSFE0RUZnUVVqWWlDdGo4blg2dHIzTHFCTURmSHU0aDdERUF3SHdZRFZSMGpCQmd3CkZvQVUxYkQ2VUUwN3cxUFBHMzZKWjJSZHJLL2lQUEl3RFFZSktvWklodmNOQVFFTEJRQURnZ0VCQUhHUWpSRUoKU3NsbFM1YkdHYkxmNGFGdVIxbDdiVGh0UTBLTFl3bkowQmtKNDBnVVZUcHdmMEFpcFRVanU1UTkxNzZaNEJWUQpleFB4U1BuRngwSGxoQXhUYmJQaWQ2Y25leVlId2I1Vi95R3I2M0p5cWcxR2pScVJ5djBLR3k0bHRnWGRKNmQyCktZaVFlT2pYN3B2SXVSMjE2T2ZNMWdnYUdOUkFzSDZwYjlEMU5ENmoyRnVlUkcrVFN1MUxzRlJXWUZsVHk0aDMKc014RElWblY3L0R1T1M5QTJCVm80alpYL0pmbzFId2ExL3BaVkJ1YVU5UHJlemNOblBkZ09TRmxVc0tBbUNjMgpCb3dES0RPbkZZMm9IZTc5ZjF6NlU1a1Buby80ZGlBT0FIQ2dPY0lPbk9ubUNOR1d3R0tLU2tGc05lYUtnVVltCjVYVkFqNVd3amNTdmhqMD0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBMWhoNzR6TUR2TXgzVGk1Q0szbDlGZjdzNXR1czcwVndHNy9yYzJqZXIxcFQ2eTFVCnlRNlU2NlVHWDU2L0h6MWdpMFpPSHhVa24xWGhtdnhmK045MCtPMGw0ejZjVjF4TUhGR20vOWhmR3NuNzJXazMKVmc2VE1MQk8rdDhxUS9QS0kwandHRUxlSXg0VUtSTUN5NDFFODFNSFMyM0w3UTB3OW1aS1N6SVhWdG9aT0h1YwpsVml2SkQrSlNDMTU2eTVVYWhuaytIWnAzTStNQzZ1a1dNZmhoMDRWekV2VW5SM1lIcnJjbEhZbE5UdUV4L0hBCjN5SkJiZlhNblNYR0krM3d4ODJIVDg0bEw1SktVYmNiSjQ0aER1YWR4N2FkTnY4djBZb3U1TWpkSzhIdUhPbksKQ29PRm9Pdmt1YklTV0ZzTFhaUGpZdFdqQURiaC9DeTVZV1VjZ3dJREFRQUJBb0lCQUFRRHo0YnlOUGE4YXR4WApkN3d5K2dxSWprN0IvZHM2MVNCZ0YvMUJFVFArb0tZL1ltQ20ybG9VN1NxcjRtK21pZ0h5bnBKc3BoUXEyeUU1CjdGN1JhZk1sRjFuTW1jZjFuaVBGMERqcUNOYUt4U05ObXRFTlV1dE4weDFYUkFha01yMDRwKy84aVFmbGo0RTUKcndxOEtuZlpyY0JYWGNTalE3RExPRWR5dUFkVDVPN0lqRjV3dENNNUw0Tlc5K0dkOFg4VnpxaTNndWlkdzUwRQphSjlndjR2Nnp3RDAydFNSZHNoUlBBbU9oVzBCVDl5bkZCNEFrdlg3b1psVHR2YU5QMVF3S0RkQXd0cmF5K2ZXCjRMcmY0ZldEWDIzNERqS01sOUN3cjhnTmNja0FoejVjN09HQ0dzeU9xc3JBTmllNFB6VWtmc0xhM1ZQQ0FsTHoKV2s1dEs0a0NnWUVBNVBMM3hzMzB2Y1NPcjhjV2R4eElCTE1pTzNnRldBb05qeEJkQ2ExNEZCYjlGaVMxK1NNSgpQTkZNTU1rT2VUMGZPamx1dVpyTGVueDIrUFZNVnhrb29kYXNmcGxtNGI3RHk1RkZzcGhUajlOOXJJdmIvWWpJCk51SDdKNkgzVEZJTzhDVE5GK1FwSmIza2pNZUhRcWM0S0cwSFBCeTlpLytvUExkOGdxZ1JCczhDZ1lFQTcyUSsKT3prYlZaUzlad25YcEZPZmQwa2hZQkkzUFlvaElockJjOEw3Q2xHZlVwdFRQdlMxSWVqUjlMNHFiM1BsSHB6bAptaG5wQXRGTFFYQXJCdFk4UmErZkxlbW1Xdzc2ZndrdlFraXlLblUvZ0s1UlVkdmdQbkxrSCtVSmJJTmwzcjgyCkRzQzB0bWJJL1QrblpObEJHdVFkQm1US3hkZ1R1MXRBZ3dIcy9BMENnWUVBMlZISDMrMmZZb0l3N3FrTHFnUXUKV0VleE5zRzJVTnM2QTVLRXZhcnJVQ2FDRllMRE9Ma0pDN0dmb0s4NERkejJ4MDI4ekhFaXRDRnd6T0FLbHFKSwo3MVBXYUZVMFV4UEF4bm9lcm1mbzZaeldyZklUMzVUMmR5SUtSSlI1S1BpN05UZTVkZlFkR3JZbE8zd3A2QnJTCk00MUtVTVQzSnV5RnhSeG1FNTkwaWdFQ2dZQjU5bHRTTnVUN00vMU8rbyszczdiaHdndFQ4OVBhOFgyeDcybXgKdlp2Q2hSVWpzK2kwZ1YycStmL0ZyZ0RXcVhnSW9hekVWd0VFbzNhd3p5SE1xT2NxSmJCMlpyeVBWZEUvV1lHUApScFFtMTNkVDZ2dVpOZWxJUjZaN3JXZWd0a3ozTC9tdGlIWkpHNUs0bTI2QURjT0NuTWRBMDZjUEp1ZmVvejM1CndNaHBIUUtCZ0E1WVpYeWE1azUrcGFRd1RaWTFKb2ZsZ2FjNlRRZ2IzOTJDWStZR1RjNHdDVVlNNWVVazg0TCsKRFM4bTYwc0liL0g3UVBOa2RBTXZobFNPMk9McFVEbGtWUnhrUXFvckxWRDRRQUc0SFBMU1VLVzI0eFNqeXpZNApncXQ4b3VmcEFIV0pMLzZpZEthUzhvVjVlc2NyV0Y5clJ6bWliUGx4V1JIZ3ZzdVJJSEhjCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==`

const bootstrapIgnition string = `
{
	"ignition":{"version":"3.1.0"},
	"storage":{
		"files":[
			{
				"path":"/opt/openshift/manifests/cvo-overrides.yaml",
				"contents":{
					"source":"data:text/plain;charset=utf-8;base64,YXBpVmVyc2lvbjogY29uZmlnLm9wZW5zaGlmdC5pby92MQpraW5kOiBDbHVzdGVyVmVyc2lvbgptZXRhZGF0YToKICBuYW1lc3BhY2U6IG9wZW5zaGlmdC1jbHVzdGVyLXZlcnNpb24KICBuYW1lOiB2ZXJzaW9uCnNwZWM6CiAgdXBzdHJlYW06IGh0dHBzOi8vYXBpLm9wZW5zaGlmdC5jb20vYXBpL3VwZ3JhZGVzX2luZm8vdjEvZ3JhcGgKICBjaGFubmVsOiBzdGFibGUtNC42CiAgY2x1c3RlcklEOiA0MTk0MGVlOC1lYzk5LTQzZGUtODc2Ni0xNzQzODFiNDkyMWQK"
				}
			}
		]
	},
	"systemd":{}
}
`
