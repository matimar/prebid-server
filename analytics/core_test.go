package analytics

import (
	"github.com/mxmCherry/openrtb"
	"github.com/prebid/prebid-server/config"
	"net/http"
	"testing"
)

func TestSampleModule(t *testing.T) {
	var count int
	am := initAnalytics(&count)
	am.LogAuctionObject(&AuctionObject{AUCTION, http.StatusOK, nil, &openrtb.BidRequest{}, &openrtb.BidResponse{}})
	if count != 1 {
		t.Errorf("PBSAnalyticsModule failed at LogAuctionObejct")
	}

	am.LogSetUIDObject(&SetUIDObject{SETUID, http.StatusOK, "bidders string", "uid", nil, true})
	if count != 2 {
		t.Errorf("PBSAnalyticsModule failed at LogSetUIDObejct")
	}

	am.LogCookieSyncObject(&CookieSyncObject{})
	if count != 3 {
		t.Errorf("PBSAnalyticsModule failed at LogCookieSyncObejct")
	}

	am.LogAmpObject(&AmpObject{})
	if count != 4 {
		t.Errorf("PBSAnalyticsModule failed at LogAmpObject")
	}
}

type sampleModule struct {
	count *int
}

func (m *sampleModule) LogAuctionObject(ao *AuctionObject) { *m.count++ }

func (m *sampleModule) LogCookieSyncObject(cso *CookieSyncObject) { *m.count++ }

func (m *sampleModule) LogSetUIDObject(so *SetUIDObject) { *m.count++ }

func (m *sampleModule) LogAmpObject(ao *AmpObject) { *m.count++ }

func initAnalytics(count *int) PBSAnalyticsModule {
	modules := make(enabledAnalytics, 0)
	modules = append(modules, &sampleModule{count})
	return &modules
}

func TestNewPBSAnalytics(t *testing.T) {
	mod := NewPBSAnalytics(&config.Analytics{File: config.FileLogs{Filename: "test"}})
	switch modType := mod.(type) {
	case enabledAnalytics:
		if len(enabledAnalytics(modType)) != 1 {
			t.Fatalf("Failed to add analytics module")
		}
	default:
		t.Fatalf("Failed to initialize analytics module")
	}
}