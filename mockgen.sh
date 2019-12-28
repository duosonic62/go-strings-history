#!/usr/bin/env bash

### use case ###
# input boundary
mockgen \
-source=internal/usecase/inputboundary/user_command.go \
-destination=internal/usecase/inputboundary/mock_inputboundary/mock_user_command.go

# output boundary
mockgen \
-source=internal/usecase/outputboundary/user_command.go \
-destination=internal/usecase/outputboundary/mock_outputboundary/mock_user_command.go

mockgen \
-source=internal/usecase/outputboundary/error.go \
-destination=internal/usecase/outputboundary/mock_outputboundary/mock_error.go

# input
mockgen \
-source=pkg/usecase/input/context.go  \
-destination=pkg/usecase/input/mock_input/mock_context.go

### domain ###
# repository
mockgen \
-source=internal/domain/repository/user.go \
-destination=internal/domain/repository/mock_repository/mock_user.go

# factory
mockgen \
-source=internal/domain/factory/user.go \
-destination=internal/domain/factory/mock_factory/mock_user.go

mockgen \
-source=internal/domain/factory/token.go \
-destination=internal/domain/factory/mock_factory/mock_token.go

mockgen \
-source=internal/domain/factory/id.go \
-destination=internal/domain/factory/mock_factory/mock_id.go