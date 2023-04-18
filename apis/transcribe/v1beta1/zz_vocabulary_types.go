/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VocabularyObservation struct {

	// ARN of the Vocabulary.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Generated download URI.
	DownloadURI *string `json:"downloadUri,omitempty" tf:"download_uri,omitempty"`

	// Name of the Vocabulary.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The language code you selected for your vocabulary.
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// - A list of terms to include in the vocabulary. Conflicts with vocabulary_file_uri
	Phrases []*string `json:"phrases,omitempty" tf:"phrases,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The Amazon S3 location (URI) of the text file that contains your custom vocabulary.
	VocabularyFileURI *string `json:"vocabularyFileUri,omitempty" tf:"vocabulary_file_uri,omitempty"`
}

type VocabularyParameters struct {

	// The language code you selected for your vocabulary.
	// +kubebuilder:validation:Optional
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// - A list of terms to include in the vocabulary. Conflicts with vocabulary_file_uri
	// +kubebuilder:validation:Optional
	Phrases []*string `json:"phrases,omitempty" tf:"phrases,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Amazon S3 location (URI) of the text file that contains your custom vocabulary.
	// +kubebuilder:validation:Optional
	VocabularyFileURI *string `json:"vocabularyFileUri,omitempty" tf:"vocabulary_file_uri,omitempty"`
}

// VocabularySpec defines the desired state of Vocabulary
type VocabularySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VocabularyParameters `json:"forProvider"`
}

// VocabularyStatus defines the observed state of Vocabulary.
type VocabularyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VocabularyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vocabulary is the Schema for the Vocabularys API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Vocabulary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.languageCode)",message="languageCode is a required parameter"
	Spec   VocabularySpec   `json:"spec"`
	Status VocabularyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VocabularyList contains a list of Vocabularys
type VocabularyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vocabulary `json:"items"`
}

// Repository type metadata.
var (
	Vocabulary_Kind             = "Vocabulary"
	Vocabulary_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vocabulary_Kind}.String()
	Vocabulary_KindAPIVersion   = Vocabulary_Kind + "." + CRDGroupVersion.String()
	Vocabulary_GroupVersionKind = CRDGroupVersion.WithKind(Vocabulary_Kind)
)

func init() {
	SchemeBuilder.Register(&Vocabulary{}, &VocabularyList{})
}
