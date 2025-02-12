// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	container "github.com/docker/docker/api/types/container"

	io "io"

	mock "github.com/stretchr/testify/mock"

	network "github.com/docker/docker/api/types/network"

	types "github.com/docker/docker/api/types"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// DockerBind is an autogenerated mock type for the DockerBind type
type DockerBind struct {
	mock.Mock
}

// ContainerCreate provides a mock function with given fields: ctx, config, hostConfig, networkingConfig, platform, containerName
func (_m *DockerBind) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *v1.Platform, containerName string) (container.ContainerCreateCreatedBody, error) {
	ret := _m.Called(ctx, config, hostConfig, networkingConfig, platform, containerName)

	var r0 container.ContainerCreateCreatedBody
	if rf, ok := ret.Get(0).(func(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, *v1.Platform, string) container.ContainerCreateCreatedBody); ok {
		r0 = rf(ctx, config, hostConfig, networkingConfig, platform, containerName)
	} else {
		r0 = ret.Get(0).(container.ContainerCreateCreatedBody)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, *v1.Platform, string) error); ok {
		r1 = rf(ctx, config, hostConfig, networkingConfig, platform, containerName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerLogs provides a mock function with given fields: ctx, _a1, options
func (_m *DockerBind) ContainerLogs(ctx context.Context, _a1 string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerLogsOptions) io.ReadCloser); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ContainerLogsOptions) error); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerRemove provides a mock function with given fields: ctx, containerID, options
func (_m *DockerBind) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error {
	ret := _m.Called(ctx, containerID, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerRemoveOptions) error); ok {
		r0 = rf(ctx, containerID, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerStart provides a mock function with given fields: ctx, containerID, options
func (_m *DockerBind) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	ret := _m.Called(ctx, containerID, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerStartOptions) error); ok {
		r0 = rf(ctx, containerID, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerWait provides a mock function with given fields: ctx, containerID, condition
func (_m *DockerBind) ContainerWait(ctx context.Context, containerID string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error) {
	ret := _m.Called(ctx, containerID, condition)

	var r0 <-chan container.ContainerWaitOKBody
	if rf, ok := ret.Get(0).(func(context.Context, string, container.WaitCondition) <-chan container.ContainerWaitOKBody); ok {
		r0 = rf(ctx, containerID, condition)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan container.ContainerWaitOKBody)
		}
	}

	var r1 <-chan error
	if rf, ok := ret.Get(1).(func(context.Context, string, container.WaitCondition) <-chan error); ok {
		r1 = rf(ctx, containerID, condition)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan error)
		}
	}

	return r0, r1
}

// ImageBuild provides a mock function with given fields: ctx, buildContext, options
func (_m *DockerBind) ImageBuild(ctx context.Context, buildContext io.Reader, options *types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	ret := _m.Called(ctx, buildContext, options)

	var r0 types.ImageBuildResponse
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, *types.ImageBuildOptions) types.ImageBuildResponse); ok {
		r0 = rf(ctx, buildContext, options)
	} else {
		r0 = ret.Get(0).(types.ImageBuildResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, io.Reader, *types.ImageBuildOptions) error); ok {
		r1 = rf(ctx, buildContext, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDockerBind interface {
	mock.TestingT
	Cleanup(func())
}

// NewDockerBind creates a new instance of DockerBind. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDockerBind(t mockConstructorTestingTNewDockerBind) *DockerBind {
	mock := &DockerBind{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
