#!/usr/bin/env bash

### use case ###
# input boundary
mockgen \
-source=internal/usecase/inputboundary/error.go \
-destination=internal/usecase/inputboundary/mock_inputboundary/mock_error.go

mockgen \
-source=internal/usecase/inputboundary/user_command.go \
-destination=internal/usecase/inputboundary/mock_inputboundary/mock_user_command.go

mockgen \
-source=internal/usecase/inputboundary/user_query.go \
-destination=internal/usecase/inputboundary/mock_inputboundary/mock_user_query.go

mockgen \
-source=internal/usecase/inputboundary/string_command.go \
-destination=internal/usecase/inputboundary/mock_inputboundary/mock_string_command.go

mockgen \
-source=internal/usecase/inputboundary/string_query.go \
-destination=internal/usecase/inputboundary/mock_inputboundary/mock_string_query.go

# output boundary
mockgen \
-source=internal/usecase/outputboundary/error.go \
-destination=internal/usecase/outputboundary/mock_outputboundary/mock_error.go

mockgen \
-source=internal/usecase/outputboundary/user_command.go \
-destination=internal/usecase/outputboundary/mock_outputboundary/mock_user_command.go

mockgen \
-source=internal/usecase/outputboundary/user_query.go \
-destination=internal/usecase/outputboundary/mock_outputboundary/mock_user_query.go

mockgen \
-source=internal/usecase/outputboundary/string_command.go \
-destination=internal/usecase/outputboundary/mock_outputboundary/mock_string_command.go

mockgen \
-source=internal/usecase/outputboundary/string_query.go \
-destination=internal/usecase/outputboundary/mock_outputboundary/mock_string_quey.go

# input
mockgen \
-source=pkg/usecase/input/context.go  \
-destination=pkg/usecase/input/mock_input/mock_context.go

### domain ###
# repository
mockgen \
-source=internal/domain/repository/user_command.go \
-destination=internal/domain/repository/mock_repository/mock_user_command.go

mockgen \
-source=internal/domain/repository/user_query.go \
-destination=internal/domain/repository/mock_repository/mock_user_query.go


mockgen \
-source=internal/domain/repository/string_command.go \
-destination=internal/domain/repository/mock_repository/mock_string_command.go

mockgen \
-source=internal/domain/repository/string_query.go \
-destination=internal/domain/repository/mock_repository/mock_string_query.go

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

mockgen \
-source=internal/domain/factory/string.go \
-destination=internal/domain/factory/mock_factory/mock_string.go

# service
mockgen \
-source=internal/domain/service/authorization.go \
-destination=internal/domain/service/mock_service/mock_authorization.go