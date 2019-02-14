// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	appsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
	imagev1 "github.com/openshift/api/image/v1"
	routev1 "github.com/openshift/api/route/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConstants) DeepCopyInto(out *AppConstants) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConstants.
func (in *AppConstants) DeepCopy() *AppConstants {
	if in == nil {
		return nil
	}
	out := new(AppConstants)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthTemplate) DeepCopyInto(out *AuthTemplate) {
	*out = *in
	in.SSO.DeepCopyInto(&out.SSO)
	out.LDAP = in.LDAP
	out.RoleMapper = in.RoleMapper
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthTemplate.
func (in *AuthTemplate) DeepCopy() *AuthTemplate {
	if in == nil {
		return nil
	}
	out := new(AuthTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildTemplate) DeepCopyInto(out *BuildTemplate) {
	*out = *in
	out.From = in.From
	out.GitSource = in.GitSource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildTemplate.
func (in *BuildTemplate) DeepCopy() *BuildTemplate {
	if in == nil {
		return nil
	}
	out := new(BuildTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonConfig) DeepCopyInto(out *CommonConfig) {
	*out = *in
	in.Auth.DeepCopyInto(&out.Auth)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonConfig.
func (in *CommonConfig) DeepCopy() *CommonConfig {
	if in == nil {
		return nil
	}
	out := new(CommonConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonKieServerSet) DeepCopyInto(out *CommonKieServerSet) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(KieAppBuildObject)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonKieServerSet.
func (in *CommonKieServerSet) DeepCopy() *CommonKieServerSet {
	if in == nil {
		return nil
	}
	out := new(CommonKieServerSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleTemplate) DeepCopyInto(out *ConsoleTemplate) {
	*out = *in
	out.SSOAuthClient = in.SSOAuthClient
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleTemplate.
func (in *ConsoleTemplate) DeepCopy() *ConsoleTemplate {
	if in == nil {
		return nil
	}
	out := new(ConsoleTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomObject) DeepCopyInto(out *CustomObject) {
	*out = *in
	if in.PersistentVolumeClaims != nil {
		in, out := &in.PersistentVolumeClaims, &out.PersistentVolumeClaims
		*out = make([]corev1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccounts != nil {
		in, out := &in.ServiceAccounts, &out.ServiceAccounts
		*out = make([]corev1.ServiceAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]corev1.Secret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]rbacv1.Role, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RoleBindings != nil {
		in, out := &in.RoleBindings, &out.RoleBindings
		*out = make([]rbacv1.RoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeploymentConfigs != nil {
		in, out := &in.DeploymentConfigs, &out.DeploymentConfigs
		*out = make([]appsv1.DeploymentConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BuildConfigs != nil {
		in, out := &in.BuildConfigs, &out.BuildConfigs
		*out = make([]buildv1.BuildConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageStreams != nil {
		in, out := &in.ImageStreams, &out.ImageStreams
		*out = make([]imagev1.ImageStream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]corev1.Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]routev1.Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomObject.
func (in *CustomObject) DeepCopy() *CustomObject {
	if in == nil {
		return nil
	}
	out := new(CustomObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvTemplate) DeepCopyInto(out *EnvTemplate) {
	*out = *in
	if in.CommonConfig != nil {
		in, out := &in.CommonConfig, &out.CommonConfig
		*out = new(CommonConfig)
		(*in).DeepCopyInto(*out)
	}
	out.Console = in.Console
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = make([]ServerTemplate, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvTemplate.
func (in *EnvTemplate) DeepCopy() *EnvTemplate {
	if in == nil {
		return nil
	}
	out := new(EnvTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	in.Console.DeepCopyInto(&out.Console)
	in.Smartrouter.DeepCopyInto(&out.Smartrouter)
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = make([]CustomObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Others != nil {
		in, out := &in.Others, &out.Others
		*out = make([]CustomObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSource) DeepCopyInto(out *GitSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSource.
func (in *GitSource) DeepCopy() *GitSource {
	if in == nil {
		return nil
	}
	out := new(GitSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KieApp) DeepCopyInto(out *KieApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KieApp.
func (in *KieApp) DeepCopy() *KieApp {
	if in == nil {
		return nil
	}
	out := new(KieApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KieApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KieAppAuthObject) DeepCopyInto(out *KieAppAuthObject) {
	*out = *in
	if in.SSO != nil {
		in, out := &in.SSO, &out.SSO
		*out = new(SSOAuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.LDAP != nil {
		in, out := &in.LDAP, &out.LDAP
		*out = new(LDAPAuthConfig)
		**out = **in
	}
	if in.RoleMapper != nil {
		in, out := &in.RoleMapper, &out.RoleMapper
		*out = new(RoleMapperAuthConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KieAppAuthObject.
func (in *KieAppAuthObject) DeepCopy() *KieAppAuthObject {
	if in == nil {
		return nil
	}
	out := new(KieAppAuthObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KieAppBuildObject) DeepCopyInto(out *KieAppBuildObject) {
	*out = *in
	out.GitSource = in.GitSource
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]WebhookSecret, len(*in))
		copy(*out, *in)
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KieAppBuildObject.
func (in *KieAppBuildObject) DeepCopy() *KieAppBuildObject {
	if in == nil {
		return nil
	}
	out := new(KieAppBuildObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KieAppList) DeepCopyInto(out *KieAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KieApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KieAppList.
func (in *KieAppList) DeepCopy() *KieAppList {
	if in == nil {
		return nil
	}
	out := new(KieAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KieAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KieAppObject) DeepCopyInto(out *KieAppObject) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KieAppObject.
func (in *KieAppObject) DeepCopy() *KieAppObject {
	if in == nil {
		return nil
	}
	out := new(KieAppObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KieAppObjects) DeepCopyInto(out *KieAppObjects) {
	*out = *in
	in.Console.DeepCopyInto(&out.Console)
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = new(CommonKieServerSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = make([]KieServerSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Smartrouter.DeepCopyInto(&out.Smartrouter)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KieAppObjects.
func (in *KieAppObjects) DeepCopy() *KieAppObjects {
	if in == nil {
		return nil
	}
	out := new(KieAppObjects)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KieAppRegistry) DeepCopyInto(out *KieAppRegistry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KieAppRegistry.
func (in *KieAppRegistry) DeepCopy() *KieAppRegistry {
	if in == nil {
		return nil
	}
	out := new(KieAppRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KieAppSpec) DeepCopyInto(out *KieAppSpec) {
	*out = *in
	out.ImageRegistry = in.ImageRegistry
	in.Objects.DeepCopyInto(&out.Objects)
	in.CommonConfig.DeepCopyInto(&out.CommonConfig)
	in.Auth.DeepCopyInto(&out.Auth)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KieAppSpec.
func (in *KieAppSpec) DeepCopy() *KieAppSpec {
	if in == nil {
		return nil
	}
	out := new(KieAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KieAppStatus) DeepCopyInto(out *KieAppStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Deployments != nil {
		in, out := &in.Deployments, &out.Deployments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KieAppStatus.
func (in *KieAppStatus) DeepCopy() *KieAppStatus {
	if in == nil {
		return nil
	}
	out := new(KieAppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KieServerSet) DeepCopyInto(out *KieServerSet) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(KieAppBuildObject)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KieServerSet.
func (in *KieServerSet) DeepCopy() *KieServerSet {
	if in == nil {
		return nil
	}
	out := new(KieServerSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LDAPAuthConfig) DeepCopyInto(out *LDAPAuthConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LDAPAuthConfig.
func (in *LDAPAuthConfig) DeepCopy() *LDAPAuthConfig {
	if in == nil {
		return nil
	}
	out := new(LDAPAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleMapperAuthConfig) DeepCopyInto(out *RoleMapperAuthConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleMapperAuthConfig.
func (in *RoleMapperAuthConfig) DeepCopy() *RoleMapperAuthConfig {
	if in == nil {
		return nil
	}
	out := new(RoleMapperAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSOAuthClient) DeepCopyInto(out *SSOAuthClient) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSOAuthClient.
func (in *SSOAuthClient) DeepCopy() *SSOAuthClient {
	if in == nil {
		return nil
	}
	out := new(SSOAuthClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSOAuthClients) DeepCopyInto(out *SSOAuthClients) {
	*out = *in
	out.Console = in.Console
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = make([]SSOAuthClient, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSOAuthClients.
func (in *SSOAuthClients) DeepCopy() *SSOAuthClients {
	if in == nil {
		return nil
	}
	out := new(SSOAuthClients)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSOAuthConfig) DeepCopyInto(out *SSOAuthConfig) {
	*out = *in
	in.Clients.DeepCopyInto(&out.Clients)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSOAuthConfig.
func (in *SSOAuthConfig) DeepCopy() *SSOAuthConfig {
	if in == nil {
		return nil
	}
	out := new(SSOAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerTemplate) DeepCopyInto(out *ServerTemplate) {
	*out = *in
	out.SSOAuthClient = in.SSOAuthClient
	out.From = in.From
	out.Build = in.Build
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerTemplate.
func (in *ServerTemplate) DeepCopy() *ServerTemplate {
	if in == nil {
		return nil
	}
	out := new(ServerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSecret) DeepCopyInto(out *WebhookSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSecret.
func (in *WebhookSecret) DeepCopy() *WebhookSecret {
	if in == nil {
		return nil
	}
	out := new(WebhookSecret)
	in.DeepCopyInto(out)
	return out
}
