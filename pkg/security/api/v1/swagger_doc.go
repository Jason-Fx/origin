package v1

// This file contains methods that can be used by the go-restful package to generate Swagger
// documentation for the object types found in 'types.go' This file is automatically generated
// by hack/update-generated-swagger-descriptions.sh and should be run after a full build of OpenShift.
// ==== DO NOT EDIT THIS FILE MANUALLY ====

var map_PodSecurityPolicyReview = map[string]string{
	"":       "PodSecurityPolicyReview checks which service accounts (not users, since that would be cluster-wide) can create the `PodSpec` in question.",
	"spec":   "spec is the PodSecurityPolicy to check.",
	"status": "status represents the current information/status for the PodSecurityPolicyReview.",
}

func (PodSecurityPolicyReview) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicyReview
}

var map_PodSecurityPolicyReviewSpec = map[string]string{
	"":                    "PodSecurityPolicyReviewSpec defines specification for PodSecurityPolicyReview",
	"podSpec":             "podSpec is the PodSpec to check. The podSpec.serviceAccountName field is used if serviceAccountNames is empty, unless the podSpec.serviceAccountName is empty, in which case \"default\" is used. If serviceAccountNames is specified, podSpec.serviceAccountName is ignored.",
	"serviceAccountNames": "serviceAccountNames is an optional set of ServiceAccounts to run the check with. If serviceAccountNames is empty, the podSpec serviceAccountName is used, unless it's empty, in which case \"default\" is used instead. If serviceAccountNames is specified, podSpec serviceAccountName is ignored.",
}

func (PodSecurityPolicyReviewSpec) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicyReviewSpec
}

var map_PodSecurityPolicyReviewStatus = map[string]string{
	"": "PodSecurityPolicyReviewStatus represents the status of PodSecurityPolicyReview.",
	"allowedServiceAccounts": "allowedServiceAccounts returns the list of service accounts in *this* namespace that have the power to create the PodSpec.",
}

func (PodSecurityPolicyReviewStatus) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicyReviewStatus
}

var map_PodSecurityPolicySelfSubjectReview = map[string]string{
	"":       "PodSecurityPolicySelfSubjectReview checks whether this user/SA tuple can create the PodSpec",
	"spec":   "spec defines specification the PodSecurityPolicySelfSubjectReview.",
	"status": "status represents the current information/status for the PodSecurityPolicySelfSubjectReview.",
}

func (PodSecurityPolicySelfSubjectReview) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySelfSubjectReview
}

var map_PodSecurityPolicySelfSubjectReviewSpec = map[string]string{
	"":        "PodSecurityPolicySelfSubjectReviewSpec contains specification for PodSecurityPolicySelfSubjectReview.",
	"podSpec": "podSpec is the PodSpec to check.",
}

func (PodSecurityPolicySelfSubjectReviewSpec) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySelfSubjectReviewSpec
}

var map_PodSecurityPolicySubjectReview = map[string]string{
	"":       "PodSecurityPolicySubjectReview checks whether a particular user/SA tuple can create the PodSpec.",
	"spec":   "spec defines specification for the PodSecurityPolicySubjectReview.",
	"status": "status represents the current information/status for the PodSecurityPolicySubjectReview.",
}

func (PodSecurityPolicySubjectReview) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySubjectReview
}

var map_PodSecurityPolicySubjectReviewSpec = map[string]string{
	"":        "PodSecurityPolicySubjectReviewSpec defines specification for PodSecurityPolicySubjectReview",
	"podSpec": "podSpec is the PodSpec to check. If podSpec.serviceAccountName is empty it will not be defaulted. If its non-empty, it will be checked.",
	"user":    "user is the user you're testing for. If you specify \"user\" but not \"group\", then is it interpreted as \"What if user were not a member of any groups. If user and groups are empty, then the check is performed using *only* the serviceAccountName in the podSpec.",
	"groups":  "groups is the groups you're testing for.",
}

func (PodSecurityPolicySubjectReviewSpec) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySubjectReviewSpec
}

var map_PodSecurityPolicySubjectReviewStatus = map[string]string{
	"":          "PodSecurityPolicySubjectReviewStatus contains information/status for PodSecurityPolicySubjectReview.",
	"allowedBy": "allowedBy is a reference to the rule that allows the PodSpec. A rule can be a SecurityContextConstraint or a PodSecurityPolicy A `nil`, indicates that it was denied.",
	"reason":    "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available.",
	"podSpec":   "podSpec is the PodSpec after the defaulting is applied.",
}

func (PodSecurityPolicySubjectReviewStatus) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySubjectReviewStatus
}

var map_ServiceAccountPodSecurityPolicyReviewStatus = map[string]string{
	"":     "ServiceAccountPodSecurityPolicyReviewStatus represents ServiceAccount name and related review status",
	"name": "name contains the allowed and the denied ServiceAccount name",
}

func (ServiceAccountPodSecurityPolicyReviewStatus) SwaggerDoc() map[string]string {
	return map_ServiceAccountPodSecurityPolicyReviewStatus
}
