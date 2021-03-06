// Copyright 2021 The Rode Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"fmt"
	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/esutil"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcErrorWithCode(log *zap.Logger, message string, err error, code codes.Code, fields ...zap.Field) error {
	if err == nil {
		log.Error(message, fields...)
		return status.Errorf(code, "%s", message)
	}

	log.Error(message, append(fields, zap.Error(err))...)
	return status.Errorf(code, "%s: %s", message, err)
}

func GrpcInternalError(log *zap.Logger, message string, err error, fields ...zap.Field) error {
	return GrpcErrorWithCode(log, message, err, codes.Internal, fields...)
}

func CheckBulkResponseErrors(response *esutil.EsBulkResponse) error {
	var bulkErrors []error
	for _, item := range response.Items {
		if item.Create != nil && item.Create.Error != nil {
			bulkErrors = append(bulkErrors, fmt.Errorf("error creating new policy version [%d] %s: %s", item.Create.Status, item.Create.Error.Type, item.Create.Error.Reason))
		} else if item.Index != nil && item.Index.Error != nil {
			bulkErrors = append(bulkErrors, fmt.Errorf("error updating policy [%d] %s: %s", item.Index.Status, item.Index.Error.Type, item.Index.Error.Reason))
		}
	}

	if len(bulkErrors) > 0 {
		return fmt.Errorf("errors: %v", bulkErrors)
	}

	return nil
}
