package templates

import (
	"fmt"

	ui "github.com/konstfish/ui/core"
	"github.com/konstfish/ui/themes/kf"
	"github.com/konstfish/whitelist-portal/portal/pkg/cache"
)

func HomePage() *ui.Page {
	page := ui.NewPage().
		SetTitle("whitelist-portal").
		AddScript("https://unpkg.com/htmx.org@2.0.4").
		AddScript("https://unpkg.com/htmx-ext-response-targets@2.0.2/response-targets.js").
		AddStyleSheet("https://ui.konst.fish/static/main.css").
		AddStyleSheet("static/main.css").
		AddScript("static/ip.js")

	content := kf.AppBody().
		AddChild(TitleBar()).
		AddChild(Form()).
		AddChild(AddressTablePlaceholder())

	page.Body.AddChild(content)

	return page
}

func TitleBar() *ui.Element {
	return kf.Panel(kf.Header1("whitelist-portal")).AddClasses("panel-adjust", "panel-header")
}

func Form() *ui.Element {
	form := kf.Form().
		AddClasses("panel", "panel-adjust", "panel-form").
		SetAttribute("hx-post", "/hx/v1/addresses").
		SetAttribute("hx-target", "#table-container").
		SetAttribute("hx-target-error", "#table-container")

	addressInput := ui.NewElement("input").
		SetAttribute("name", "address").
		SetAttribute("id", "address").
		SetAttribute("maxlength", "15")

	addressGroup := kf.GroupClass("form-opt address",
		ui.NewElement("label").SetContent("Address"),
		addressInput,
	)

	expiryInput := ui.NewElement("input").
		SetAttribute("type", "number").
		SetAttribute("name", "expiry").
		SetAttribute("min", "1").
		SetAttribute("value", "90").
		SetAttribute("list", "expiryList")

	expiryGroup := kf.GroupClass("form-opt expiry",
		ui.NewElement("label").SetContent("Expiry (Days)"),
		expiryInput,
	)

	firstRow := kf.GroupClass("form-row",
		addressGroup,
		expiryGroup,
	)

	descriptionInput := ui.NewElement("input").
		SetAttribute("name", "description").
		SetAttribute("id", "description").
		SetAttribute("maxlength", "20")

	descriptionGroup := kf.GroupClass("form-opt description",
		ui.NewElement("label").SetContent("Description"),
		descriptionInput,
	)

	submitButton := ui.NewElement("button").
		AddClass("btn-submit").
		SetContent("Add")

	secondRow := kf.GroupClass("form-row",
		descriptionGroup,
		submitButton,
	)

	form.AddChild(firstRow)
	form.AddChild(secondRow)

	return form
}

func AddressTablePlaceholder() *ui.Element {
	placeholder := kf.Group(kf.Spinner("Loading...")).
		AddClasses("panel", "panel-adjust", "panel-table").
		SetId("table-container").
		SetAttribute("hx-get", "/hx/v1/addresses").
		SetAttribute("hx-trigger", "load").
		SetAttribute("hx-target-error", "#table-container")

	return placeholder
}

func AddressTableEntry(item cache.AddressListEntry) *ui.Element {
	entry := kf.Group().
		AddClasses("panel-adjust", "panel-address").
		AddChild(kf.GroupClass("address-info", kf.Header3(item.Address), kf.Paragraph(item.Description))).
		AddChild(kf.GroupClass("address-button", kf.Paragraph(ttlToReadable(item.TTL)), kf.Button("x").AddClasses("btn-delete").SetAttribute("hx-delete", "/hx/v1/addresses/"+item.Address).SetAttribute("hx-target", "#table-container")))

	return entry
}

func ttlToReadable(ttl int64) string {
	return fmt.Sprintf("%dd", ttl/24)
}

func ErrorMessage(message string) *ui.Element {
	return kf.GroupClass("info-warning", kf.Paragraph(message))
}
