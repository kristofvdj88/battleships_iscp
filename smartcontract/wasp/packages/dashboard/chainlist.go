package dashboard

import (
	"net/http"

	"github.com/iotaledger/wasp/packages/registry"
	"github.com/iotaledger/wasp/plugins/chains"
	"github.com/labstack/echo/v4"
)

func initChainList(e *echo.Echo, r renderer) Tab {
	route := e.GET("/chains", handleChainList)
	route.Name = "chainList"

	r[route.Path] = makeTemplate(e, tplChainList)

	return Tab{
		Path:  route.Path,
		Title: "Chains",
		Href:  route.Path,
	}
}

func handleChainList(c echo.Context) error {
	chains, err := fetchChains()
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, c.Path(), &ChainListTemplateParams{
		BaseTemplateParams: BaseParams(c),
		Chains:             chains,
	})
}

func fetchChains() ([]*ChainOverview, error) {
	crs, err := registry.GetChainRecords()
	if err != nil {
		return nil, err
	}
	r := make([]*ChainOverview, len(crs))
	for i, cr := range crs {
		info, err := fetchRootInfo(chains.GetChain(cr.ChainID))
		r[i] = &ChainOverview{
			ChainRecord: cr,
			RootInfo:    info,
			Error:       err,
		}
	}
	return r, nil
}

type ChainListTemplateParams struct {
	BaseTemplateParams
	Chains []*ChainOverview
}

type ChainOverview struct {
	ChainRecord *registry.ChainRecord
	RootInfo    RootInfo
	Error       error
}

const tplChainList = `
{{define "title"}}Chains{{end}}

{{define "body"}}
<div class="card fluid">
	<h2 class="section">Chains</h2>
	<table>
		<thead>
			<tr>
				<th>ID</th>
				<th>Description</th>
				<th>#Nodes</th>
				<th>#Contracts</th>
				<th>Active?</th>
			</tr>
		</thead>
		<tbody>
			{{range $_, $c := .Chains}}
				{{ $id := $c.ChainRecord.ChainID }}
				<tr>
					<td data-label="ID">{{ if not $c.Error }}<a href="{{ uri "chain" $id }}"><tt>{{ $id }}</tt></a>{{ else }}<tt>{{ $id }}</tt>{{ end }}</td>
					<td data-label="Description">{{ trim 50 $c.RootInfo.Description }}
						{{- if $c.Error }}<div class="card fluid error">{{ $c.Error }}</div>{{ end }}</td>
					<td data-label="#Nodes">{{if not $c.Error}}<tt>{{ len $c.ChainRecord.CommitteeNodes }}</tt>{{ end }}</td>
					<td data-label="#Contracts">{{if not $c.Error}}<tt>{{ len $c.RootInfo.Contracts }}</tt>{{ end }}</td>
					<td data-label="Active?">{{ if $c.ChainRecord.Active }} yes {{ else }} no {{ end }}</td>
				</tr>
			{{end}}
		</tbody>
	</table>
</div>
{{end}}
`
