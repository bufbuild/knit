// @generated by protoc-gen-connect-es v0.8.6 with parameter "target=ts"
// @generated from file buf/starwars/starship/v1/starship.proto (package buf.starwars.starship.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { StarshipsRequest, StarshipsResponse } from "./starship_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service buf.starwars.starship.v1.StarshipService
 */
export const StarshipService = {
  typeName: "buf.starwars.starship.v1.StarshipService",
  methods: {
    /**
     * @generated from rpc buf.starwars.starship.v1.StarshipService.GetStarships
     */
    getStarships: {
      name: "GetStarships",
      I: StarshipsRequest,
      O: StarshipsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

