// Automatically generated by MockGen. DO NOT EDIT!
// Source: kubevirt.go

package kubecli

import (
	io "io"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	discovery "k8s.io/client-go/discovery"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1"
	v1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	v1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	v10 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	v1beta10 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	v11 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	v1beta11 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	v12 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	v2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	v13 "k8s.io/client-go/kubernetes/typed/batch/v1"
	v1beta12 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	v2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1"
	v1beta13 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	v14 "k8s.io/client-go/kubernetes/typed/core/v1"
	v1beta14 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	v15 "k8s.io/client-go/kubernetes/typed/networking/v1"
	v1beta15 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	v16 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	v1alpha10 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	v1beta16 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	v1alpha11 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	v1alpha12 "k8s.io/client-go/kubernetes/typed/settings/v1alpha1"
	v17 "k8s.io/client-go/kubernetes/typed/storage/v1"
	v1beta17 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	rest "k8s.io/client-go/rest"

	v18 "kubevirt.io/kubevirt/pkg/api/v1"
)

// Mock of KubevirtClient interface
type MockKubevirtClient struct {
	ctrl     *gomock.Controller
	recorder *_MockKubevirtClientRecorder
}

// Recorder for MockKubevirtClient (not exported)
type _MockKubevirtClientRecorder struct {
	mock *MockKubevirtClient
}

func NewMockKubevirtClient(ctrl *gomock.Controller) *MockKubevirtClient {
	mock := &MockKubevirtClient{ctrl: ctrl}
	mock.recorder = &_MockKubevirtClientRecorder{mock}
	return mock
}

func (_m *MockKubevirtClient) EXPECT() *_MockKubevirtClientRecorder {
	return _m.recorder
}

func (_m *MockKubevirtClient) VM(namespace string) VMInterface {
	ret := _m.ctrl.Call(_m, "VM", namespace)
	ret0, _ := ret[0].(VMInterface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) VM(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "VM", arg0)
}

func (_m *MockKubevirtClient) ReplicaSet(namespace string) ReplicaSetInterface {
	ret := _m.ctrl.Call(_m, "ReplicaSet", namespace)
	ret0, _ := ret[0].(ReplicaSetInterface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) ReplicaSet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReplicaSet", arg0)
}

func (_m *MockKubevirtClient) RestClient() *rest.RESTClient {
	ret := _m.ctrl.Call(_m, "RestClient")
	ret0, _ := ret[0].(*rest.RESTClient)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) RestClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RestClient")
}

func (_m *MockKubevirtClient) Discovery() discovery.DiscoveryInterface {
	ret := _m.ctrl.Call(_m, "Discovery")
	ret0, _ := ret[0].(discovery.DiscoveryInterface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Discovery() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Discovery")
}

func (_m *MockKubevirtClient) AdmissionregistrationV1alpha1() v1alpha1.AdmissionregistrationV1alpha1Interface {
	ret := _m.ctrl.Call(_m, "AdmissionregistrationV1alpha1")
	ret0, _ := ret[0].(v1alpha1.AdmissionregistrationV1alpha1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) AdmissionregistrationV1alpha1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AdmissionregistrationV1alpha1")
}

func (_m *MockKubevirtClient) Admissionregistration() v1alpha1.AdmissionregistrationV1alpha1Interface {
	ret := _m.ctrl.Call(_m, "Admissionregistration")
	ret0, _ := ret[0].(v1alpha1.AdmissionregistrationV1alpha1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Admissionregistration() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Admissionregistration")
}

func (_m *MockKubevirtClient) AppsV1beta1() v1beta1.AppsV1beta1Interface {
	ret := _m.ctrl.Call(_m, "AppsV1beta1")
	ret0, _ := ret[0].(v1beta1.AppsV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) AppsV1beta1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AppsV1beta1")
}

func (_m *MockKubevirtClient) AppsV1beta2() v1beta2.AppsV1beta2Interface {
	ret := _m.ctrl.Call(_m, "AppsV1beta2")
	ret0, _ := ret[0].(v1beta2.AppsV1beta2Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) AppsV1beta2() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AppsV1beta2")
}

func (_m *MockKubevirtClient) Apps() v1beta2.AppsV1beta2Interface {
	ret := _m.ctrl.Call(_m, "Apps")
	ret0, _ := ret[0].(v1beta2.AppsV1beta2Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Apps() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Apps")
}

func (_m *MockKubevirtClient) AuthenticationV1() v10.AuthenticationV1Interface {
	ret := _m.ctrl.Call(_m, "AuthenticationV1")
	ret0, _ := ret[0].(v10.AuthenticationV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) AuthenticationV1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AuthenticationV1")
}

func (_m *MockKubevirtClient) Authentication() v10.AuthenticationV1Interface {
	ret := _m.ctrl.Call(_m, "Authentication")
	ret0, _ := ret[0].(v10.AuthenticationV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Authentication() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Authentication")
}

func (_m *MockKubevirtClient) AuthenticationV1beta1() v1beta10.AuthenticationV1beta1Interface {
	ret := _m.ctrl.Call(_m, "AuthenticationV1beta1")
	ret0, _ := ret[0].(v1beta10.AuthenticationV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) AuthenticationV1beta1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AuthenticationV1beta1")
}

func (_m *MockKubevirtClient) AuthorizationV1() v11.AuthorizationV1Interface {
	ret := _m.ctrl.Call(_m, "AuthorizationV1")
	ret0, _ := ret[0].(v11.AuthorizationV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) AuthorizationV1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AuthorizationV1")
}

func (_m *MockKubevirtClient) Authorization() v11.AuthorizationV1Interface {
	ret := _m.ctrl.Call(_m, "Authorization")
	ret0, _ := ret[0].(v11.AuthorizationV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Authorization() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Authorization")
}

func (_m *MockKubevirtClient) AuthorizationV1beta1() v1beta11.AuthorizationV1beta1Interface {
	ret := _m.ctrl.Call(_m, "AuthorizationV1beta1")
	ret0, _ := ret[0].(v1beta11.AuthorizationV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) AuthorizationV1beta1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AuthorizationV1beta1")
}

func (_m *MockKubevirtClient) AutoscalingV1() v12.AutoscalingV1Interface {
	ret := _m.ctrl.Call(_m, "AutoscalingV1")
	ret0, _ := ret[0].(v12.AutoscalingV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) AutoscalingV1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AutoscalingV1")
}

func (_m *MockKubevirtClient) Autoscaling() v12.AutoscalingV1Interface {
	ret := _m.ctrl.Call(_m, "Autoscaling")
	ret0, _ := ret[0].(v12.AutoscalingV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Autoscaling() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Autoscaling")
}

func (_m *MockKubevirtClient) AutoscalingV2beta1() v2beta1.AutoscalingV2beta1Interface {
	ret := _m.ctrl.Call(_m, "AutoscalingV2beta1")
	ret0, _ := ret[0].(v2beta1.AutoscalingV2beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) AutoscalingV2beta1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AutoscalingV2beta1")
}

func (_m *MockKubevirtClient) BatchV1() v13.BatchV1Interface {
	ret := _m.ctrl.Call(_m, "BatchV1")
	ret0, _ := ret[0].(v13.BatchV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) BatchV1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchV1")
}

func (_m *MockKubevirtClient) Batch() v13.BatchV1Interface {
	ret := _m.ctrl.Call(_m, "Batch")
	ret0, _ := ret[0].(v13.BatchV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Batch() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Batch")
}

func (_m *MockKubevirtClient) BatchV1beta1() v1beta12.BatchV1beta1Interface {
	ret := _m.ctrl.Call(_m, "BatchV1beta1")
	ret0, _ := ret[0].(v1beta12.BatchV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) BatchV1beta1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchV1beta1")
}

func (_m *MockKubevirtClient) BatchV2alpha1() v2alpha1.BatchV2alpha1Interface {
	ret := _m.ctrl.Call(_m, "BatchV2alpha1")
	ret0, _ := ret[0].(v2alpha1.BatchV2alpha1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) BatchV2alpha1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchV2alpha1")
}

func (_m *MockKubevirtClient) CertificatesV1beta1() v1beta13.CertificatesV1beta1Interface {
	ret := _m.ctrl.Call(_m, "CertificatesV1beta1")
	ret0, _ := ret[0].(v1beta13.CertificatesV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) CertificatesV1beta1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CertificatesV1beta1")
}

func (_m *MockKubevirtClient) Certificates() v1beta13.CertificatesV1beta1Interface {
	ret := _m.ctrl.Call(_m, "Certificates")
	ret0, _ := ret[0].(v1beta13.CertificatesV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Certificates() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Certificates")
}

func (_m *MockKubevirtClient) CoreV1() v14.CoreV1Interface {
	ret := _m.ctrl.Call(_m, "CoreV1")
	ret0, _ := ret[0].(v14.CoreV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) CoreV1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CoreV1")
}

func (_m *MockKubevirtClient) Core() v14.CoreV1Interface {
	ret := _m.ctrl.Call(_m, "Core")
	ret0, _ := ret[0].(v14.CoreV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Core() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Core")
}

func (_m *MockKubevirtClient) ExtensionsV1beta1() v1beta14.ExtensionsV1beta1Interface {
	ret := _m.ctrl.Call(_m, "ExtensionsV1beta1")
	ret0, _ := ret[0].(v1beta14.ExtensionsV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) ExtensionsV1beta1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExtensionsV1beta1")
}

func (_m *MockKubevirtClient) Extensions() v1beta14.ExtensionsV1beta1Interface {
	ret := _m.ctrl.Call(_m, "Extensions")
	ret0, _ := ret[0].(v1beta14.ExtensionsV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Extensions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Extensions")
}

func (_m *MockKubevirtClient) NetworkingV1() v15.NetworkingV1Interface {
	ret := _m.ctrl.Call(_m, "NetworkingV1")
	ret0, _ := ret[0].(v15.NetworkingV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) NetworkingV1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NetworkingV1")
}

func (_m *MockKubevirtClient) Networking() v15.NetworkingV1Interface {
	ret := _m.ctrl.Call(_m, "Networking")
	ret0, _ := ret[0].(v15.NetworkingV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Networking() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Networking")
}

func (_m *MockKubevirtClient) PolicyV1beta1() v1beta15.PolicyV1beta1Interface {
	ret := _m.ctrl.Call(_m, "PolicyV1beta1")
	ret0, _ := ret[0].(v1beta15.PolicyV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) PolicyV1beta1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PolicyV1beta1")
}

func (_m *MockKubevirtClient) Policy() v1beta15.PolicyV1beta1Interface {
	ret := _m.ctrl.Call(_m, "Policy")
	ret0, _ := ret[0].(v1beta15.PolicyV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Policy() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Policy")
}

func (_m *MockKubevirtClient) RbacV1() v16.RbacV1Interface {
	ret := _m.ctrl.Call(_m, "RbacV1")
	ret0, _ := ret[0].(v16.RbacV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) RbacV1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RbacV1")
}

func (_m *MockKubevirtClient) Rbac() v16.RbacV1Interface {
	ret := _m.ctrl.Call(_m, "Rbac")
	ret0, _ := ret[0].(v16.RbacV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Rbac() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rbac")
}

func (_m *MockKubevirtClient) RbacV1beta1() v1beta16.RbacV1beta1Interface {
	ret := _m.ctrl.Call(_m, "RbacV1beta1")
	ret0, _ := ret[0].(v1beta16.RbacV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) RbacV1beta1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RbacV1beta1")
}

func (_m *MockKubevirtClient) RbacV1alpha1() v1alpha10.RbacV1alpha1Interface {
	ret := _m.ctrl.Call(_m, "RbacV1alpha1")
	ret0, _ := ret[0].(v1alpha10.RbacV1alpha1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) RbacV1alpha1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RbacV1alpha1")
}

func (_m *MockKubevirtClient) SchedulingV1alpha1() v1alpha11.SchedulingV1alpha1Interface {
	ret := _m.ctrl.Call(_m, "SchedulingV1alpha1")
	ret0, _ := ret[0].(v1alpha11.SchedulingV1alpha1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) SchedulingV1alpha1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SchedulingV1alpha1")
}

func (_m *MockKubevirtClient) Scheduling() v1alpha11.SchedulingV1alpha1Interface {
	ret := _m.ctrl.Call(_m, "Scheduling")
	ret0, _ := ret[0].(v1alpha11.SchedulingV1alpha1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Scheduling() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Scheduling")
}

func (_m *MockKubevirtClient) SettingsV1alpha1() v1alpha12.SettingsV1alpha1Interface {
	ret := _m.ctrl.Call(_m, "SettingsV1alpha1")
	ret0, _ := ret[0].(v1alpha12.SettingsV1alpha1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) SettingsV1alpha1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SettingsV1alpha1")
}

func (_m *MockKubevirtClient) Settings() v1alpha12.SettingsV1alpha1Interface {
	ret := _m.ctrl.Call(_m, "Settings")
	ret0, _ := ret[0].(v1alpha12.SettingsV1alpha1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Settings() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Settings")
}

func (_m *MockKubevirtClient) StorageV1beta1() v1beta17.StorageV1beta1Interface {
	ret := _m.ctrl.Call(_m, "StorageV1beta1")
	ret0, _ := ret[0].(v1beta17.StorageV1beta1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) StorageV1beta1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StorageV1beta1")
}

func (_m *MockKubevirtClient) StorageV1() v17.StorageV1Interface {
	ret := _m.ctrl.Call(_m, "StorageV1")
	ret0, _ := ret[0].(v17.StorageV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) StorageV1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StorageV1")
}

func (_m *MockKubevirtClient) Storage() v17.StorageV1Interface {
	ret := _m.ctrl.Call(_m, "Storage")
	ret0, _ := ret[0].(v17.StorageV1Interface)
	return ret0
}

func (_mr *_MockKubevirtClientRecorder) Storage() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Storage")
}

// Mock of VMInterface interface
type MockVMInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockVMInterfaceRecorder
}

// Recorder for MockVMInterface (not exported)
type _MockVMInterfaceRecorder struct {
	mock *MockVMInterface
}

func NewMockVMInterface(ctrl *gomock.Controller) *MockVMInterface {
	mock := &MockVMInterface{ctrl: ctrl}
	mock.recorder = &_MockVMInterfaceRecorder{mock}
	return mock
}

func (_m *MockVMInterface) EXPECT() *_MockVMInterfaceRecorder {
	return _m.recorder
}

func (_m *MockVMInterface) Get(name string, options v1.GetOptions) (*v18.VirtualMachine, error) {
	ret := _m.ctrl.Call(_m, "Get", name, options)
	ret0, _ := ret[0].(*v18.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMInterfaceRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}

func (_m *MockVMInterface) List(opts v1.ListOptions) (*v18.VirtualMachineList, error) {
	ret := _m.ctrl.Call(_m, "List", opts)
	ret0, _ := ret[0].(*v18.VirtualMachineList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMInterfaceRecorder) List(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List", arg0)
}

func (_m *MockVMInterface) Create(_param0 *v18.VirtualMachine) (*v18.VirtualMachine, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0)
	ret0, _ := ret[0].(*v18.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMInterfaceRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0)
}

func (_m *MockVMInterface) Update(_param0 *v18.VirtualMachine) (*v18.VirtualMachine, error) {
	ret := _m.ctrl.Call(_m, "Update", _param0)
	ret0, _ := ret[0].(*v18.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMInterfaceRecorder) Update(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0)
}

func (_m *MockVMInterface) Delete(name string, options *v1.DeleteOptions) error {
	ret := _m.ctrl.Call(_m, "Delete", name, options)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVMInterfaceRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1)
}

func (_m *MockVMInterface) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v18.VirtualMachine, error) {
	_s := []interface{}{name, pt, data}
	for _, _x := range subresources {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Patch", _s...)
	ret0, _ := ret[0].(*v18.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMInterfaceRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Patch", _s...)
}

func (_m *MockVMInterface) SerialConsole(name string, device string, in io.Reader, out io.Writer) error {
	ret := _m.ctrl.Call(_m, "SerialConsole", name, device, in, out)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVMInterfaceRecorder) SerialConsole(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SerialConsole", arg0, arg1, arg2, arg3)
}

func (_m *MockVMInterface) VNC(name string, in io.Reader, out io.Writer) error {
	ret := _m.ctrl.Call(_m, "VNC", name, in, out)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVMInterfaceRecorder) VNC(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "VNC", arg0, arg1, arg2)
}

// Mock of ReplicaSetInterface interface
type MockReplicaSetInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockReplicaSetInterfaceRecorder
}

// Recorder for MockReplicaSetInterface (not exported)
type _MockReplicaSetInterfaceRecorder struct {
	mock *MockReplicaSetInterface
}

func NewMockReplicaSetInterface(ctrl *gomock.Controller) *MockReplicaSetInterface {
	mock := &MockReplicaSetInterface{ctrl: ctrl}
	mock.recorder = &_MockReplicaSetInterfaceRecorder{mock}
	return mock
}

func (_m *MockReplicaSetInterface) EXPECT() *_MockReplicaSetInterfaceRecorder {
	return _m.recorder
}

func (_m *MockReplicaSetInterface) Get(name string, options v1.GetOptions) (*v18.VirtualMachineReplicaSet, error) {
	ret := _m.ctrl.Call(_m, "Get", name, options)
	ret0, _ := ret[0].(*v18.VirtualMachineReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockReplicaSetInterfaceRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}

func (_m *MockReplicaSetInterface) List(opts v1.ListOptions) (*v18.VirtualMachineReplicaSetList, error) {
	ret := _m.ctrl.Call(_m, "List", opts)
	ret0, _ := ret[0].(*v18.VirtualMachineReplicaSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockReplicaSetInterfaceRecorder) List(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List", arg0)
}

func (_m *MockReplicaSetInterface) Create(_param0 *v18.VirtualMachineReplicaSet) (*v18.VirtualMachineReplicaSet, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0)
	ret0, _ := ret[0].(*v18.VirtualMachineReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockReplicaSetInterfaceRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0)
}

func (_m *MockReplicaSetInterface) Update(_param0 *v18.VirtualMachineReplicaSet) (*v18.VirtualMachineReplicaSet, error) {
	ret := _m.ctrl.Call(_m, "Update", _param0)
	ret0, _ := ret[0].(*v18.VirtualMachineReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockReplicaSetInterfaceRecorder) Update(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0)
}

func (_m *MockReplicaSetInterface) Delete(name string, options *v1.DeleteOptions) error {
	ret := _m.ctrl.Call(_m, "Delete", name, options)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockReplicaSetInterfaceRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1)
}

// Mock of VMPresetInterface interface
type MockVMPresetInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockVMPresetInterfaceRecorder
}

// Recorder for MockVMPresetInterface (not exported)
type _MockVMPresetInterfaceRecorder struct {
	mock *MockVMPresetInterface
}

func NewMockVMPresetInterface(ctrl *gomock.Controller) *MockVMPresetInterface {
	mock := &MockVMPresetInterface{ctrl: ctrl}
	mock.recorder = &_MockVMPresetInterfaceRecorder{mock}
	return mock
}

func (_m *MockVMPresetInterface) EXPECT() *_MockVMPresetInterfaceRecorder {
	return _m.recorder
}

func (_m *MockVMPresetInterface) Get(name string, options v1.GetOptions) (*v18.VirtualMachinePreset, error) {
	ret := _m.ctrl.Call(_m, "Get", name, options)
	ret0, _ := ret[0].(*v18.VirtualMachinePreset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMPresetInterfaceRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}

func (_m *MockVMPresetInterface) List(opts v1.ListOptions) (*v18.VirtualMachinePresetList, error) {
	ret := _m.ctrl.Call(_m, "List", opts)
	ret0, _ := ret[0].(*v18.VirtualMachinePresetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMPresetInterfaceRecorder) List(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List", arg0)
}

func (_m *MockVMPresetInterface) Create(_param0 *v18.VirtualMachinePreset) (*v18.VirtualMachinePreset, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0)
	ret0, _ := ret[0].(*v18.VirtualMachinePreset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMPresetInterfaceRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0)
}

func (_m *MockVMPresetInterface) Update(_param0 *v18.VirtualMachinePreset) (*v18.VirtualMachinePreset, error) {
	ret := _m.ctrl.Call(_m, "Update", _param0)
	ret0, _ := ret[0].(*v18.VirtualMachinePreset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMPresetInterfaceRecorder) Update(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0)
}

func (_m *MockVMPresetInterface) Delete(name string, options *v1.DeleteOptions) error {
	ret := _m.ctrl.Call(_m, "Delete", name, options)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVMPresetInterfaceRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1)
}

func (_m *MockVMPresetInterface) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v18.VirtualMachinePreset, error) {
	_s := []interface{}{name, pt, data}
	for _, _x := range subresources {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Patch", _s...)
	ret0, _ := ret[0].(*v18.VirtualMachinePreset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVMPresetInterfaceRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Patch", _s...)
}
