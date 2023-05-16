// @generated by protoc-gen-es v1.2.0 with parameter "target=ts"
// @generated from file buf/starwars/film/v1/film.proto (package buf.starwars.film.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message buf.starwars.film.v1.Film
 */
export class Film extends Message<Film> {
  /**
   * @generated from field: string film_id = 1;
   */
  filmId = "";

  /**
   * @generated from field: uint64 episode_number = 2;
   */
  episodeNumber = protoInt64.zero;

  /**
   * @generated from field: string title = 3;
   */
  title = "";

  /**
   * @generated from field: string opening_text = 4;
   */
  openingText = "";

  /**
   * @generated from field: repeated string directors = 5;
   */
  directors: string[] = [];

  /**
   * @generated from field: repeated string producers = 6;
   */
  producers: string[] = [];

  /**
   * @generated from field: repeated string character_ids = 7;
   */
  characterIds: string[] = [];

  /**
   * @generated from field: repeated string planet_ids = 8;
   */
  planetIds: string[] = [];

  /**
   * @generated from field: repeated string starship_ids = 9;
   */
  starshipIds: string[] = [];

  /**
   * @generated from field: repeated string vehicle_ids = 10;
   */
  vehicleIds: string[] = [];

  /**
   * @generated from field: repeated string species_ids = 11;
   */
  speciesIds: string[] = [];

  constructor(data?: PartialMessage<Film>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "buf.starwars.film.v1.Film";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "film_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "episode_number", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "opening_text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "directors", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "producers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 7, name: "character_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 8, name: "planet_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 9, name: "starship_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 10, name: "vehicle_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 11, name: "species_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Film {
    return new Film().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Film {
    return new Film().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Film {
    return new Film().fromJsonString(jsonString, options);
  }

  static equals(a: Film | PlainMessage<Film> | undefined, b: Film | PlainMessage<Film> | undefined): boolean {
    return proto3.util.equals(Film, a, b);
  }
}

/**
 * @generated from message buf.starwars.film.v1.FilmsRequest
 */
export class FilmsRequest extends Message<FilmsRequest> {
  /**
   * @generated from field: repeated string film_ids = 1;
   */
  filmIds: string[] = [];

  /**
   * @generated from field: uint32 limit = 2;
   */
  limit = 0;

  constructor(data?: PartialMessage<FilmsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "buf.starwars.film.v1.FilmsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "film_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "limit", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FilmsRequest {
    return new FilmsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FilmsRequest {
    return new FilmsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FilmsRequest {
    return new FilmsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: FilmsRequest | PlainMessage<FilmsRequest> | undefined, b: FilmsRequest | PlainMessage<FilmsRequest> | undefined): boolean {
    return proto3.util.equals(FilmsRequest, a, b);
  }
}

/**
 * @generated from message buf.starwars.film.v1.FilmsResponse
 */
export class FilmsResponse extends Message<FilmsResponse> {
  /**
   * @generated from field: repeated buf.starwars.film.v1.Film films = 1;
   */
  films: Film[] = [];

  constructor(data?: PartialMessage<FilmsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "buf.starwars.film.v1.FilmsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "films", kind: "message", T: Film, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FilmsResponse {
    return new FilmsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FilmsResponse {
    return new FilmsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FilmsResponse {
    return new FilmsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: FilmsResponse | PlainMessage<FilmsResponse> | undefined, b: FilmsResponse | PlainMessage<FilmsResponse> | undefined): boolean {
    return proto3.util.equals(FilmsResponse, a, b);
  }
}

