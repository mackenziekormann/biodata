package AlphaFold

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
    "net/http"
)

// Creating a data structure mirroring the output of an AlphaFold request
type ProteinData struct {
    EntryID                string   `json:"entryId"`
    Gene                   string   `json:"gene"`
    UniprotAccession       string   `json:"uniprotAccession"`
    UniprotID              string   `json:"uniprotId"`
    UniprotDescription     string   `json:"uniprotDescription"`
    TaxID                  int      `json:"taxId"`
    OrganismScientificName string   `json:"organismScientificName"`
    UniprotStart           int      `json:"uniprotStart"`
    UniprotEnd             int      `json:"uniprotEnd"`
    UniprotSequence        string   `json:"uniprotSequence"`
    ModelCreatedDate       string   `json:"modelCreatedDate"`
    LatestVersion          int      `json:"latestVersion"`
    AllVersions            []int    `json:"allVersions"`
    IsReviewed             bool     `json:"isReviewed"`
    IsReferenceProteome    bool     `json:"isReferenceProteome"`
    CIFUrl                 string   `json:"cifUrl"`
    BCIFUrl                string   `json:"bcifUrl"`
    PDBUrl                 string   `json:"pdbUrl"`
    PAEImageUrl            string   `json:"paeImageUrl"`
    PAEDocUrl              string   `json:"paeDocUrl"`
}


// Fetch protein structure data from AlphaFold API
func FetchAlphaFoldData(proteinID string) (*ProteinData, error) {
	url := fmt.Sprintf("https://alphafold.ebi.ac.uk/api/prediction/%s", proteinID)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var data ProteinData
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, err
    }

    return &data, nil
}

// Processes protein data from AlphaFold to prepare for visualization
func ProcessProteinData(rawData ProteinData) (ProcessedData, error) {
	//TO-DO
}