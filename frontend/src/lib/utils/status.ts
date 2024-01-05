export let statusOK = 200;
export let statusCreated = 201;
export let statusNoContent = 204;
export let statusConflict = 409;
export let statusBadRequest = 400;
export let statusUnauthorized = 401;
export let statusNotFound = 404;
export let statusInternalServerError = 500;

export default {
    OK: statusOK,
    Created: statusCreated,
    NoContent: statusNoContent,
    Conflict: statusConflict,
    BadRequest: statusBadRequest,
    Unauthorized: statusUnauthorized,
    NotFound: statusNotFound,
    InternalServerError: statusInternalServerError,
};