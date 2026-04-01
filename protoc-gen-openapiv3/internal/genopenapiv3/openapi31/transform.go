package openapi31

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv3/internal/genopenapiv3/model"
)

// TransformDocument converts a canonical Document to OpenAPI 3.1.0 output format.
func TransformDocument(doc *model.Document) *Document {
	if doc == nil {
		return nil
	}

	return &Document{
		OpenAPI:      "3.1.0",
		Info:         transformInfo(doc.Info),
		Servers:      transformServers(doc.Servers),
		Paths:        transformPaths(doc.Paths),
		Components:   transformComponents(doc.Components),
		Security:     transformSecurityRequirements(doc.Security),
		Tags:         transformTags(doc.Tags),
		ExternalDocs: transformExternalDocs(doc.ExternalDocs),
	}
}

func transformInfo(info *model.Info) *Info {
	if info == nil {
		return nil
	}
	return &Info{
		Title:          info.Title,
		Summary:        info.Summary,
		Description:    info.Description,
		TermsOfService: info.TermsOfService,
		Contact:        transformContact(info.Contact),
		License:        transformLicense(info.License),
		Version:        info.Version,
	}
}

func transformContact(c *model.Contact) *Contact {
	if c == nil {
		return nil
	}
	return &Contact{
		Name:  c.Name,
		URL:   c.URL,
		Email: c.Email,
	}
}

func transformLicense(l *model.License) *License {
	if l == nil {
		return nil
	}
	return &License{
		Name:       l.Name,
		Identifier: l.Identifier,
		URL:        l.URL,
	}
}

func transformServers(servers []*model.Server) []*Server {
	if len(servers) == 0 {
		return nil
	}
	result := make([]*Server, len(servers))
	for i, s := range servers {
		result[i] = transformServer(s)
	}
	return result
}

func transformServer(s *model.Server) *Server {
	if s == nil {
		return nil
	}
	var vars map[string]*ServerVariable
	if len(s.Variables) > 0 {
		vars = make(map[string]*ServerVariable)
		for name, v := range s.Variables {
			vars[name] = &ServerVariable{
				Enum:        v.Enum,
				Default:     v.Default,
				Description: v.Description,
			}
		}
	}
	return &Server{
		URL:         s.URL,
		Description: s.Description,
		Variables:   vars,
	}
}

func transformPaths(paths *model.Paths) *Paths {
	if paths == nil || len(paths.Items) == 0 {
		return nil
	}
	result := &Paths{
		Items: make(map[string]*PathItem),
	}
	for path, item := range paths.Items {
		result.Items[path] = transformPathItem(item)
	}
	return result
}

func transformPathItem(item *model.PathItem) *PathItem {
	if item == nil {
		return nil
	}
	return &PathItem{
		Ref:         item.Ref,
		Summary:     item.Summary,
		Description: item.Description,
		Get:         transformOperation(item.Get),
		Put:         transformOperation(item.Put),
		Post:        transformOperation(item.Post),
		Delete:      transformOperation(item.Delete),
		Options:     transformOperation(item.Options),
		Head:        transformOperation(item.Head),
		Patch:       transformOperation(item.Patch),
		Trace:       transformOperation(item.Trace),
		Servers:     transformServers(item.Servers),
		Parameters:  transformParameterOrRefs(item.Parameters),
	}
}

func transformOperation(op *model.Operation) *Operation {
	if op == nil {
		return nil
	}
	return &Operation{
		Tags:         op.Tags,
		Summary:      op.Summary,
		Description:  op.Description,
		ExternalDocs: transformExternalDocs(op.ExternalDocs),
		OperationID:  op.OperationID,
		Parameters:   transformParameterOrRefs(op.Parameters),
		RequestBody:  transformRequestBodyOrRef(op.RequestBody),
		Responses:    transformResponses(op.Responses),
		Callbacks:    transformCallbacks(op.Callbacks),
		Deprecated:   op.Deprecated,
		Security:     transformSecurityRequirements(op.Security),
		Servers:      transformServers(op.Servers),
	}
}

func transformExternalDocs(ed *model.ExternalDocs) *ExternalDocs {
	if ed == nil {
		return nil
	}
	return &ExternalDocs{
		Description: ed.Description,
		URL:         ed.URL,
	}
}

func transformParameterOrRefs(params []*model.ParameterOrRef) []*ParameterOrRef {
	if len(params) == 0 {
		return nil
	}
	result := make([]*ParameterOrRef, len(params))
	for i, p := range params {
		result[i] = transformParameterOrRef(p)
	}
	return result
}

func transformParameterOrRef(p *model.ParameterOrRef) *ParameterOrRef {
	if p == nil {
		return nil
	}
	if p.Ref != "" {
		return &ParameterOrRef{Ref: p.Ref}
	}
	if p.Value == nil {
		return nil
	}
	return &ParameterOrRef{
		Value: &Parameter{
			Name:            p.Value.Name,
			In:              p.Value.In,
			Description:     p.Value.Description,
			Required:        p.Value.Required,
			Deprecated:      p.Value.Deprecated,
			AllowEmptyValue: p.Value.AllowEmptyValue,
			Style:           p.Value.Style,
			Explode:         p.Value.Explode,
			AllowReserved:   p.Value.AllowReserved,
			Schema:          transformSchemaOrRef(p.Value.Schema),
			Examples:        transformExamples(p.Value.Examples),
			Content:         transformContent(p.Value.Content),
		},
	}
}

func transformRequestBodyOrRef(rb *model.RequestBodyOrRef) *RequestBodyOrRef {
	if rb == nil {
		return nil
	}
	if rb.Ref != "" {
		return &RequestBodyOrRef{Ref: rb.Ref}
	}
	if rb.Value == nil {
		return nil
	}
	return &RequestBodyOrRef{
		Value: &RequestBody{
			Description: rb.Value.Description,
			Content:     transformContent(rb.Value.Content),
			Required:    rb.Value.Required,
		},
	}
}

func transformContent(content map[string]*model.MediaType) map[string]*MediaType {
	if len(content) == 0 {
		return nil
	}
	result := make(map[string]*MediaType)
	for mt, media := range content {
		result[mt] = transformMediaType(media)
	}
	return result
}

func transformMediaType(mt *model.MediaType) *MediaType {
	if mt == nil {
		return nil
	}
	return &MediaType{
		Schema:   transformSchemaOrRef(mt.Schema),
		Examples: transformExamples(mt.Examples),
		Encoding: transformEncoding(mt.Encoding),
	}
}

func transformEncoding(encoding map[string]*model.Encoding) map[string]*Encoding {
	if len(encoding) == 0 {
		return nil
	}
	result := make(map[string]*Encoding)
	for name, enc := range encoding {
		result[name] = &Encoding{
			ContentType:   enc.ContentType,
			Headers:       transformHeaderOrRefs(enc.Headers),
			Style:         enc.Style,
			Explode:       enc.Explode,
			AllowReserved: enc.AllowReserved,
		}
	}
	return result
}

func transformResponses(r *model.Responses) *Responses {
	if r == nil {
		return nil
	}
	result := &Responses{
		Default: transformResponseOrRef(r.Default),
		Codes:   make(map[string]*ResponseOrRef),
	}
	for code, resp := range r.Codes {
		result.Codes[code] = transformResponseOrRef(resp)
	}
	return result
}

func transformResponseOrRef(r *model.ResponseOrRef) *ResponseOrRef {
	if r == nil {
		return nil
	}
	if r.Ref != "" {
		return &ResponseOrRef{Ref: r.Ref}
	}
	if r.Value == nil {
		return nil
	}
	return &ResponseOrRef{
		Value: &Response{
			Description: r.Value.Description,
			Headers:     transformHeaderOrRefs(r.Value.Headers),
			Content:     transformContent(r.Value.Content),
			Links:       transformLinkOrRefs(r.Value.Links),
		},
	}
}

func transformHeaderOrRefs(headers map[string]*model.HeaderOrRef) map[string]*HeaderOrRef {
	if len(headers) == 0 {
		return nil
	}
	result := make(map[string]*HeaderOrRef)
	for name, h := range headers {
		result[name] = transformHeaderOrRef(h)
	}
	return result
}

func transformHeaderOrRef(h *model.HeaderOrRef) *HeaderOrRef {
	if h == nil {
		return nil
	}
	if h.Ref != "" {
		return &HeaderOrRef{Ref: h.Ref}
	}
	if h.Value == nil {
		return nil
	}
	return &HeaderOrRef{
		Value: &Header{
			Description:     h.Value.Description,
			Required:        h.Value.Required,
			Deprecated:      h.Value.Deprecated,
			AllowEmptyValue: h.Value.AllowEmptyValue,
			Style:           h.Value.Style,
			Explode:         h.Value.Explode,
			AllowReserved:   h.Value.AllowReserved,
			Schema:          transformSchemaOrRef(h.Value.Schema),
			Examples:        transformExamples(h.Value.Examples),
			Content:         transformContent(h.Value.Content),
		},
	}
}

func transformLinkOrRefs(links map[string]*model.LinkOrRef) map[string]*LinkOrRef {
	if len(links) == 0 {
		return nil
	}
	result := make(map[string]*LinkOrRef)
	for name, l := range links {
		result[name] = transformLinkOrRef(l)
	}
	return result
}

func transformLinkOrRef(l *model.LinkOrRef) *LinkOrRef {
	if l == nil {
		return nil
	}
	if l.Ref != "" {
		return &LinkOrRef{Ref: l.Ref}
	}
	if l.Value == nil {
		return nil
	}
	return &LinkOrRef{
		Value: &Link{
			OperationRef: l.Value.OperationRef,
			OperationID:  l.Value.OperationID,
			Parameters:   l.Value.Parameters,
			RequestBody:  l.Value.RequestBody,
			Description:  l.Value.Description,
			Server:       transformServer(l.Value.Server),
		},
	}
}

func transformCallbacks(callbacks map[string]*model.CallbackOrRef) map[string]*CallbackOrRef {
	if len(callbacks) == 0 {
		return nil
	}
	result := make(map[string]*CallbackOrRef)
	for name, cb := range callbacks {
		result[name] = transformCallbackOrRef(cb)
	}
	return result
}

func transformCallbackOrRef(cb *model.CallbackOrRef) *CallbackOrRef {
	if cb == nil {
		return nil
	}
	if cb.Ref != "" {
		return &CallbackOrRef{Ref: cb.Ref}
	}
	if cb.Value == nil {
		return nil
	}
	pathItems := make(map[string]*PathItem)
	for expr, item := range cb.Value {
		pathItems[expr] = transformPathItem(item)
	}
	return &CallbackOrRef{Value: pathItems}
}

// transformSchemaOrRef transforms a canonical SchemaOrRef to 3.1.0 format.
// In OpenAPI 3.1.0, $ref can have sibling summary/description.
func transformSchemaOrRef(s *model.SchemaOrRef) *SchemaOrRef {
	if s == nil {
		return nil
	}
	if s.Ref != "" {
		// 3.1.0 allows $ref with sibling summary/description
		return &SchemaOrRef{
			Ref:         s.Ref,
			Summary:     s.Summary,
			Description: s.Description,
		}
	}
	if s.Value == nil {
		return nil
	}
	return &SchemaOrRef{
		Value: transformSchema(s.Value),
	}
}

// transformSchema transforms a canonical Schema to 3.1.0 format.
// Key difference: nullable is expressed as type array ["type", "null"] in 3.1.0.
func transformSchema(s *model.Schema) *Schema {
	if s == nil {
		return nil
	}

	// Handle nullable via type array in 3.1.0
	var typeVal any
	if s.Type != "" {
		if s.IsNullable {
			// 3.1.0 style: type array with null
			typeVal = []string{s.Type, "null"}
		} else {
			typeVal = s.Type
		}
	}

	return &Schema{
		Type:                 typeVal,
		Format:               s.Format,
		Title:                s.Title,
		Description:          s.Description,
		Default:              s.Default,
		Examples:             s.Examples,
		Deprecated:           s.Deprecated,
		ReadOnly:             s.ReadOnly,
		WriteOnly:            s.WriteOnly,
		ExternalDocs:         transformExternalDocs(s.ExternalDocs),
		MultipleOf:           s.MultipleOf,
		Minimum:              s.Minimum,
		Maximum:              s.Maximum,
		ExclusiveMinimum:     s.ExclusiveMinimum,
		ExclusiveMaximum:     s.ExclusiveMaximum,
		MinLength:            s.MinLength,
		MaxLength:            s.MaxLength,
		Pattern:              s.Pattern,
		MinItems:             s.MinItems,
		MaxItems:             s.MaxItems,
		UniqueItems:          s.UniqueItems,
		Items:                transformSchemaOrRef(s.Items),
		MinProperties:        s.MinProperties,
		MaxProperties:        s.MaxProperties,
		Required:             s.Required,
		Properties:           transformSchemaProperties(s.Properties),
		AdditionalProperties: transformAdditionalProperties(s.AdditionalProperties),
		AllOf:                transformSchemaOrRefs(s.AllOf),
		AnyOf:                transformSchemaOrRefs(s.AnyOf),
		OneOf:                transformSchemaOrRefs(s.OneOf),
		Not:                  transformSchemaOrRef(s.Not),
		Discriminator:        transformDiscriminator(s.Discriminator),
		Enum:                 s.Enum,
	}
}

func transformSchemaProperties(props map[string]*model.SchemaOrRef) map[string]*SchemaOrRef {
	if len(props) == 0 {
		return nil
	}
	result := make(map[string]*SchemaOrRef)
	for name, prop := range props {
		result[name] = transformSchemaOrRef(prop)
	}
	return result
}

func transformSchemaOrRefs(schemas []*model.SchemaOrRef) []*SchemaOrRef {
	if len(schemas) == 0 {
		return nil
	}
	result := make([]*SchemaOrRef, len(schemas))
	for i, s := range schemas {
		result[i] = transformSchemaOrRef(s)
	}
	return result
}

func transformAdditionalProperties(ap *model.AdditionalProperties) *AdditionalProperties {
	if ap == nil {
		return nil
	}
	return &AdditionalProperties{
		Allowed: ap.Allowed,
		Schema:  transformSchemaOrRef(ap.Schema),
	}
}

func transformDiscriminator(d *model.Discriminator) *Discriminator {
	if d == nil {
		return nil
	}
	return &Discriminator{
		PropertyName: d.PropertyName,
		Mapping:      d.Mapping,
	}
}

func transformExamples(examples map[string]*model.Example) map[string]*Example {
	if len(examples) == 0 {
		return nil
	}
	result := make(map[string]*Example)
	for name, ex := range examples {
		if ex == nil {
			continue
		}
		result[name] = &Example{
			Ref:           ex.Ref,
			Summary:       ex.Summary,
			Description:   ex.Description,
			Value:         ex.Value,
			ExternalValue: ex.ExternalValue,
		}
	}
	return result
}

func transformComponents(c *model.Components) *Components {
	if c == nil {
		return nil
	}
	return &Components{
		Schemas:         transformComponentSchemas(c.Schemas),
		Responses:       transformComponentResponses(c.Responses),
		Parameters:      transformComponentParameters(c.Parameters),
		Examples:        transformComponentExamples(c.Examples),
		RequestBodies:   transformComponentRequestBodies(c.RequestBodies),
		Headers:         transformComponentHeaders(c.Headers),
		SecuritySchemes: transformComponentSecuritySchemes(c.SecuritySchemes),
		Links:           transformComponentLinks(c.Links),
		Callbacks:       transformComponentCallbacks(c.Callbacks),
	}
}

func transformComponentSchemas(schemas map[string]*model.SchemaOrRef) map[string]*SchemaOrRef {
	if len(schemas) == 0 {
		return nil
	}
	result := make(map[string]*SchemaOrRef)
	for name, s := range schemas {
		result[name] = transformSchemaOrRef(s)
	}
	return result
}

func transformComponentResponses(responses map[string]*model.ResponseOrRef) map[string]*ResponseOrRef {
	if len(responses) == 0 {
		return nil
	}
	result := make(map[string]*ResponseOrRef)
	for name, r := range responses {
		result[name] = transformResponseOrRef(r)
	}
	return result
}

func transformComponentParameters(params map[string]*model.ParameterOrRef) map[string]*ParameterOrRef {
	if len(params) == 0 {
		return nil
	}
	result := make(map[string]*ParameterOrRef)
	for name, p := range params {
		result[name] = transformParameterOrRef(p)
	}
	return result
}

func transformComponentExamples(examples map[string]*model.Example) map[string]*Example {
	if len(examples) == 0 {
		return nil
	}
	result := make(map[string]*Example)
	for name, ex := range examples {
		if ex == nil {
			continue
		}
		result[name] = &Example{
			Ref:           ex.Ref,
			Summary:       ex.Summary,
			Description:   ex.Description,
			Value:         ex.Value,
			ExternalValue: ex.ExternalValue,
		}
	}
	return result
}

func transformComponentRequestBodies(bodies map[string]*model.RequestBodyOrRef) map[string]*RequestBodyOrRef {
	if len(bodies) == 0 {
		return nil
	}
	result := make(map[string]*RequestBodyOrRef)
	for name, rb := range bodies {
		result[name] = transformRequestBodyOrRef(rb)
	}
	return result
}

func transformComponentHeaders(headers map[string]*model.HeaderOrRef) map[string]*HeaderOrRef {
	if len(headers) == 0 {
		return nil
	}
	result := make(map[string]*HeaderOrRef)
	for name, h := range headers {
		result[name] = transformHeaderOrRef(h)
	}
	return result
}

func transformComponentSecuritySchemes(schemes map[string]*model.SecuritySchemeOrRef) map[string]*SecuritySchemeOrRef {
	if len(schemes) == 0 {
		return nil
	}
	result := make(map[string]*SecuritySchemeOrRef)
	for name, ss := range schemes {
		result[name] = transformSecuritySchemeOrRef(ss)
	}
	return result
}

func transformSecuritySchemeOrRef(ss *model.SecuritySchemeOrRef) *SecuritySchemeOrRef {
	if ss == nil {
		return nil
	}
	if ss.Ref != "" {
		return &SecuritySchemeOrRef{Ref: ss.Ref}
	}
	if ss.Value == nil {
		return nil
	}
	return &SecuritySchemeOrRef{
		Value: &SecurityScheme{
			Type:             ss.Value.Type,
			Description:      ss.Value.Description,
			Name:             ss.Value.Name,
			In:               ss.Value.In,
			Scheme:           ss.Value.Scheme,
			BearerFormat:     ss.Value.BearerFormat,
			Flows:            transformOAuthFlows(ss.Value.Flows),
			OpenIDConnectURL: ss.Value.OpenIDConnectURL,
		},
	}
}

func transformOAuthFlows(flows *model.OAuthFlows) *OAuthFlows {
	if flows == nil {
		return nil
	}
	return &OAuthFlows{
		Implicit:          transformOAuthFlow(flows.Implicit),
		Password:          transformOAuthFlow(flows.Password),
		ClientCredentials: transformOAuthFlow(flows.ClientCredentials),
		AuthorizationCode: transformOAuthFlow(flows.AuthorizationCode),
	}
}

func transformOAuthFlow(flow *model.OAuthFlow) *OAuthFlow {
	if flow == nil {
		return nil
	}
	return &OAuthFlow{
		AuthorizationURL: flow.AuthorizationURL,
		TokenURL:         flow.TokenURL,
		RefreshURL:       flow.RefreshURL,
		Scopes:           flow.Scopes,
	}
}

func transformComponentLinks(links map[string]*model.LinkOrRef) map[string]*LinkOrRef {
	if len(links) == 0 {
		return nil
	}
	result := make(map[string]*LinkOrRef)
	for name, l := range links {
		result[name] = transformLinkOrRef(l)
	}
	return result
}

func transformComponentCallbacks(callbacks map[string]*model.CallbackOrRef) map[string]*CallbackOrRef {
	if len(callbacks) == 0 {
		return nil
	}
	result := make(map[string]*CallbackOrRef)
	for name, cb := range callbacks {
		result[name] = transformCallbackOrRef(cb)
	}
	return result
}

func transformSecurityRequirements(reqs []model.SecurityRequirement) []SecurityRequirement {
	if len(reqs) == 0 {
		return nil
	}
	result := make([]SecurityRequirement, len(reqs))
	for i, req := range reqs {
		result[i] = SecurityRequirement(req)
	}
	return result
}

func transformTags(tags []*model.Tag) []*Tag {
	if len(tags) == 0 {
		return nil
	}
	result := make([]*Tag, len(tags))
	for i, t := range tags {
		result[i] = &Tag{
			Name:         t.Name,
			Description:  t.Description,
			ExternalDocs: transformExternalDocs(t.ExternalDocs),
		}
	}
	return result
}
