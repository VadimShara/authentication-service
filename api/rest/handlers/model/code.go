package model

// @Enum CodesEnum
type CodesEnum struct {
	CodesEnum string `enum:"login_success,user_created,jwt_recieved,username_is_taken,wrong_credentials,invalid_jwt,invalid_id,invalid_request_body,invalid_header,invalid_query_params,id_is_required,request_body_is_required,header_is_required,auth_header_is_required,star_received,stars_received,star_created,star_updated,star_deleted,star_not_found,movie_received,movies_received,movie_created,movie_updated,movie_deleted,movie_not_found,general_unauthorized,general_access_denied,general_internal,general_bad_request_error,general_unsupported_method,general_forbidden"`
}