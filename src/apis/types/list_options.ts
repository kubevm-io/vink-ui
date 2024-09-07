// @generated by protobuf-ts 2.9.4
// @generated from protobuf file "types/list_options.proto" (package "vink.kubevm.io.apis.types", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { NamespaceName } from "./namespace_name";
/**
 * @generated from protobuf message vink.kubevm.io.apis.types.ListOptions
 */
export interface ListOptions {
    /**
     * @generated from protobuf field: string label_selector = 1;
     */
    labelSelector: string;
    /**
     * @generated from protobuf field: string field_selector = 2;
     */
    fieldSelector: string;
    /**
     * @generated from protobuf field: int32 limit = 3;
     */
    limit: number;
    /**
     * @generated from protobuf field: string continue = 4;
     */
    continue: string;
    /**
     * @generated from protobuf field: repeated vink.kubevm.io.apis.types.NamespaceName namespace_names = 5;
     */
    namespaceNames: NamespaceName[];
    /**
     * @generated from protobuf field: bool watch = 6;
     */
    watch: boolean;
}
// @generated message type with reflection information, may provide speed optimized methods
class ListOptions$Type extends MessageType<ListOptions> {
    constructor() {
        super("vink.kubevm.io.apis.types.ListOptions", [
            { no: 1, name: "label_selector", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "field_selector", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "limit", kind: "scalar", T: 5 /*ScalarType.INT32*/, options: { "validate.rules": { int32: { gte: 0 } } } },
            { no: 4, name: "continue", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "namespace_names", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => NamespaceName },
            { no: 6, name: "watch", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<ListOptions>): ListOptions {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.labelSelector = "";
        message.fieldSelector = "";
        message.limit = 0;
        message.continue = "";
        message.namespaceNames = [];
        message.watch = false;
        if (value !== undefined)
            reflectionMergePartial<ListOptions>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListOptions): ListOptions {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string label_selector */ 1:
                    message.labelSelector = reader.string();
                    break;
                case /* string field_selector */ 2:
                    message.fieldSelector = reader.string();
                    break;
                case /* int32 limit */ 3:
                    message.limit = reader.int32();
                    break;
                case /* string continue */ 4:
                    message.continue = reader.string();
                    break;
                case /* repeated vink.kubevm.io.apis.types.NamespaceName namespace_names */ 5:
                    message.namespaceNames.push(NamespaceName.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* bool watch */ 6:
                    message.watch = reader.bool();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ListOptions, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string label_selector = 1; */
        if (message.labelSelector !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.labelSelector);
        /* string field_selector = 2; */
        if (message.fieldSelector !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.fieldSelector);
        /* int32 limit = 3; */
        if (message.limit !== 0)
            writer.tag(3, WireType.Varint).int32(message.limit);
        /* string continue = 4; */
        if (message.continue !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.continue);
        /* repeated vink.kubevm.io.apis.types.NamespaceName namespace_names = 5; */
        for (let i = 0; i < message.namespaceNames.length; i++)
            NamespaceName.internalBinaryWrite(message.namespaceNames[i], writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* bool watch = 6; */
        if (message.watch !== false)
            writer.tag(6, WireType.Varint).bool(message.watch);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message vink.kubevm.io.apis.types.ListOptions
 */
export const ListOptions = new ListOptions$Type();
