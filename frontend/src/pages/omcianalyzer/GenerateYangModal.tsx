import { useState, useEffect } from "react";
import axios from "axios";
import { Button, Input, Select, SelectItem, Modal, ModalContent, ModalHeader, ModalBody, ModalFooter, Chip, Textarea } from "@heroui/react";
import { API_BASE } from "@/config/api";

const baseUrl = API_BASE + "/api/omcianalyzer";

interface GenerateYangModalProps {
  isOpen: boolean;
  onClose: () => void;
  requestKey: string;
  onuName: string;
}

interface GenerateResult {
  onuName?: string;
  xml?: string;
  services?: number;
  valid?: boolean;
  gaps?: string[];
  notes?: string[];
  validationErrors?: string[];
  validationWarnings?: string[];
  summary?: string;
  board?: string;
  error?: string;
}

export default function GenerateYangModal({ isOpen, onClose, requestKey, onuName }: GenerateYangModalProps) {
  const [boards, setBoards] = useState<string[]>([]);
  const [board, setBoard] = useState("");
  const [olName, setOlName] = useState("");
  const [vendor, setVendor] = useState("");
  const [chassisName, setChassisName] = useState("");
  const [veipComponent, setVeipComponent] = useState("");
  const [busy, setBusy] = useState(false);
  const [result, setResult] = useState<GenerateResult | null>(null);
  const [error, setError] = useState<string | null>(null);

  // The boards on offer come from the backend, which reads them off the YANG
  // sets actually shipped -- a hard-coded list here would drift from them.
  useEffect(() => {
    if (!isOpen) return;
    setResult(null);
    setError(null);
    setOlName(onuName);
    axios
      .get(`${baseUrl}/yangboards`)
      .then((resp) => {
        const list: string[] = resp.data?.boards ?? [];
        setBoards(list);
        setBoard((current) => current || list[0] || "");
      })
      .catch((err) => setError(err?.message ?? "cannot list YANG boards"));
  }, [isOpen, onuName]);

  const generate = async () => {
    setBusy(true);
    setError(null);
    setResult(null);
    try {
      const resp = await axios.post(`${baseUrl}/generate`, {
        requestKey,
        onuName,
        oltOnuName: olName,
        board,
        vendor,
        chassisName,
        veipComponent,
      });
      setResult(resp.data);
    } catch (err: any) {
      // A config that failed to compile still comes back with a body saying
      // why, which is more useful than the status code.
      const body = err?.response?.data;
      if (body) {
        setResult(body);
      } else {
        setError(err?.message ?? "request failed");
      }
    } finally {
      setBusy(false);
    }
  };

  const download = () => {
    if (!result?.xml) return;
    const blob = new Blob([result.xml], { type: "application/xml" });
    const url = URL.createObjectURL(blob);
    const anchor = document.createElement("a");
    anchor.href = url;
    anchor.download = `${olName || onuName}-${(result.board ?? board).toLowerCase()}.xml`;
    anchor.click();
    URL.revokeObjectURL(url);
  };

  return (
    <Modal isOpen={isOpen} onOpenChange={onClose} size="4xl" scrollBehavior="inside">
      <ModalContent>
        <ModalHeader className="flex flex-col gap-1">
          <span>Generate LightSpan Configuration</span>
          <span className="text-sm font-normal text-default-500">
            ONU {onuName}
          </span>
        </ModalHeader>
        <ModalBody>
          <div className="grid grid-cols-2 gap-3">
            <Select
              label="LT board"
              description="Selects the YANG set. The mounted tree is served by the LT's confd, and the platform modules differ between boards."
              selectedKeys={board ? [board] : []}
              onChange={(e) => setBoard(e.target.value)}
              isRequired
            >
              {boards.map((b) => (
                <SelectItem key={b}>{b}</SelectItem>
              ))}
            </Select>
            <Input
              label="ONU name on the OLT"
              description="How the target OLT refers to this ONU, which need not be what the log called it. OMCI does not carry it, and without it the config cannot be addressed."
              value={olName}
              onValueChange={setOlName}
              isRequired
            />
            <Input
              label="Vendor OUI"
              description="Selects vendor YANGMAP overrides, e.g. ALCL. Left empty it is reported as a gap."
              value={vendor}
              onValueChange={setVendor}
            />
            <Input
              label="Chassis name (optional)"
              description="Reuse a chassis the OLT already holds; a second one is rejected outright."
              value={chassisName}
              onValueChange={setChassisName}
            />
            <Input
              label="VEIP component (optional)"
              description="The virtual-UNI component a VEIP rides. Only an HGU reports one."
              value={veipComponent}
              onValueChange={setVeipComponent}
            />
          </div>

          {error && <div className="text-sm text-red-600">Error: {error}</div>}

          {result && (
            <div className="mt-2 flex flex-col gap-2">
              {result.error ? (
                <div className="text-sm text-red-600">{result.error}</div>
              ) : (
                <>
                  <div className="flex gap-2 items-center flex-wrap">
                    <Chip color={result.valid ? "success" : "danger"} variant="flat">
                      {result.valid ? "schema-valid" : "invalid"}
                    </Chip>
                    <Chip variant="flat">{result.services ?? 0} service(s)</Chip>
                    <Chip color={result.gaps?.length ? "warning" : "default"} variant="flat">
                      {result.gaps?.length ?? 0} gap(s)
                    </Chip>
                    <span className="text-sm text-default-500">{result.summary}</span>
                  </div>

                  {/* A gap means the document is schema-correct but incomplete,
                      so it is shown next to the XML rather than hidden. */}
                  {!!result.gaps?.length && (
                    <div className="text-sm">
                      <div className="font-medium">Gaps -- these need filling in by hand:</div>
                      <ul className="list-disc ml-5">
                        {result.gaps.map((gap, i) => (
                          <li key={i} className="text-amber-700">{gap}</li>
                        ))}
                      </ul>
                    </div>
                  )}
                  {!!result.notes?.length && (
                    <div className="text-sm">
                      <div className="font-medium">Notes:</div>
                      <ul className="list-disc ml-5">
                        {result.notes.map((note, i) => (
                          <li key={i} className="text-default-600">{note}</li>
                        ))}
                      </ul>
                    </div>
                  )}
                  {!!result.validationErrors?.length && (
                    <div className="text-sm">
                      <div className="font-medium text-red-600">Schema errors:</div>
                      <ul className="list-disc ml-5">
                        {result.validationErrors.map((e, i) => (
                          <li key={i} className="text-red-600">{e}</li>
                        ))}
                      </ul>
                    </div>
                  )}

                  {result.xml && (
                    <Textarea
                      label="edit-config"
                      value={result.xml}
                      minRows={14}
                      maxRows={14}
                      isReadOnly
                      classNames={{ input: "font-mono text-xs" }}
                    />
                  )}
                </>
              )}
            </div>
          )}
        </ModalBody>
        <ModalFooter>
          <Button variant="light" onPress={onClose}>Close</Button>
          {result?.xml && (
            <Button color="default" onPress={download}>Download XML</Button>
          )}
          <Button
            color="primary"
            onPress={generate}
            isLoading={busy}
            isDisabled={!board || !olName}
          >
            Generate
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
}
